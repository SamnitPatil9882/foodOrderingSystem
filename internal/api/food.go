package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/food"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/gorilla/mux"
)

func GetFoodListHandler(foodSvc food.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		respones, err := foodSvc.GetFoodList(ctx)
		if err != nil {
			fmt.Println("error occured in Getting food list in handler: " + err.Error())
		}
		json.NewEncoder(w).Encode(respones)
	}
}

func GetFoodListByCategoryHandler(foodSvc food.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var id int
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["category_id"])
		if err != nil {
			fmt.Println("error occured in parsing int in GetFoodListByCategory: " + err.Error())
		}
		ctx := r.Context()
		respones, err := foodSvc.GetFoodListByCategory(ctx, id)
		if err != nil {
			fmt.Println("error occured in Getting food list in handler: " + err.Error())
		}
		json.NewEncoder(w).Encode(respones)
	}
}

func CreateFoodItemHandler(foodSvc food.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.FoodCreateRequest
		json.NewDecoder(r.Body).Decode(&req)
		response, err := foodSvc.CreateFoodItem(ctx, req)
		if err != nil {
			fmt.Println("Error in create food Handler: " + err.Error())
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}

func UpdateFoodItemHandler(foodSvc food.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.Food
		json.NewDecoder(r.Body).Decode(&req)
		response, err := foodSvc.UpdateFoodItem(ctx, req)
		if err != nil {
			fmt.Println("Error in update food Handler: " + err.Error())
			w.WriteHeader(404)
			return
		}
		json.NewEncoder(w).Encode(response)

	}
}
