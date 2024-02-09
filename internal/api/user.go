package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/user"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
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
		err = userSvc.Signup(ctx, sgnUpReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
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
			w.Write([]byte(err.Error()))
		}
		err = userSvc.Login(ctx, loginReq)
		if err != nil {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func GetUsersHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		response, err := userSvc.GetUsers(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}
func GetUserHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println("error occured in parsing int in GetFoodListByCategory: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		respones, err := userSvc.GetUser(ctx, id)
		if err != nil {
			fmt.Println("error occured in Getting food list in handler: " + err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(respones)

	}
}
