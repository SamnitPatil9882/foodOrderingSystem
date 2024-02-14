package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/user"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func SignUpHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sgnUpReq := dto.UserSignUpRequest{}

		err := json.NewDecoder(r.Body).Decode(&sgnUpReq)
		if err != nil {
			log.Fatal(err)
			w.Write([]byte(err.Error()))
		}

		err = validateCreateUserReq(&sgnUpReq)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := userSvc.Signup(ctx, sgnUpReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		token, err := GenerateJWT(resp)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		w.Header().Set("Authorization", "Bearer"+token)
		w.WriteHeader(http.StatusAccepted)
	}
}

func LoginHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		loginReq := dto.UserLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		valres := isValidEmail(loginReq.Email) && isValidPassword(loginReq.Password)
		if !valres {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidEmail}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := userSvc.Login(ctx, loginReq)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.InvalidCredentials}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		log.Printf("log in resp: %v", resp)
		token, err := GenerateJWT(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		w.Header().Set("Authorization", "Bearer"+token)
		w.WriteHeader(http.StatusAccepted)
	}
}

func GetUsersHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := userSvc.GetUsers(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}
func GetUserHandler(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("error occured in parsing int in GetFoodListByCategory: " + err.Error())
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		ctx := r.Context()
		respones, err := userSvc.GetUser(ctx, id)
		if err != nil {
			log.Println("error occured in Getting food list in handler: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(respones)

	}
}

func UpdateUserHandler(userSvc user.Service)  http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userInfo := dto.UpdateUserInfo{}
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = validateUpdateUserInfo(userInfo)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := userSvc.UpdateUser(ctx, userInfo, getUserID(r))
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}
func GenerateJWT(user dto.UserLoginResponse) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(time.Hour * 1)

	// Create claims
	claims := jwt.MapClaims{
		"userid": user.ID,
		"role":   user.Role,
		"exp":    expirationTime.Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(constants.JWTKEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
