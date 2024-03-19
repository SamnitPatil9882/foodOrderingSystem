package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/gorilla/mux"
)

func GetFoodListHandler(foodSvc food.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		respones, err := foodSvc.GetFoodList(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = json.NewEncoder(w).Encode(respones)
		if err != nil {
			log.Println("Handler error: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
	}
}

func GetFoodListByCategoryHandler(foodSvc food.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["category_id"])
		if err != nil {
			log.Println("error occured in parsing int in GetFoodListByCategory: " + err.Error())
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "enter valid data", http.StatusBadRequest)
			return
		}
		if id <= 0 {
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "enter valid data", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		respones, err := foodSvc.GetFoodListByCategory(ctx, id)
		if err != nil {
			log.Println("error occured in Getting food list in handler: " + err.Error())
			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(respones)
	}
}

func GetFoodInfoByIDHandler(foodSvc food.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["food_id"])
		if err != nil {
			log.Println("error occured in parsing int in GetFoodListByCategory: " + err.Error())
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "enter valid data", http.StatusBadRequest)
			return
		}
		if id <= 0 {
			// w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "enter valid data", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		respones, err := foodSvc.GetFoodInfoByID(ctx, id)
		if err != nil {
			log.Println("error occured in Getting food list in handler: " + err.Error())
			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(respones)
	}
}
func CreateFoodItemHandler(foodSvc food.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.FoodCreateRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, internal.InternalServerError, http.StatusInternalServerError)
			return
		}
		err = validateFoodCreateReq(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		response, err := foodSvc.CreateFoodItem(ctx, req)
		if err != nil {
			log.Println("Error in create food Handler: " + err.Error())
			// w.WriteHeader(404)
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}

func UpdateFoodItemHandler(foodSvc food.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.Food
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err.Error())
			// w.Write([]byte("enter valid attributes"))
			// w.WriteHeader(http.StatusBadRequest)
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = validateUpdateFoodReq(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		response, err := foodSvc.UpdateFoodItem(ctx, req)
		if err != nil {
			log.Println("Error in update food Handler: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}
