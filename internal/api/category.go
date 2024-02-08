package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/category"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/gorilla/mux"
)

func GetCategoriesHandler(categorySvc category.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// ctx := r.Context()

		respone, err := categorySvc.GetCategories(ctx)
		if err != nil {
			fmt.Println("Handler error: " + err.Error())
		}
		json.NewEncoder(w).Encode(respone)
		// w.WriteHeader(http.StatusOK)

	}

}
func GetCategoryHandler(categorySvc category.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// ctx := r.Context()

		params := mux.Vars(r)
		category_id, err := strconv.ParseInt(params["category_id"], 10, 64)
		fmt.Println(category_id)
		if err != nil {
			fmt.Println("error occured in parsing string to int")
		}
		respone, err := categorySvc.GetCategory(ctx, category_id)
		if err != nil {
			fmt.Println("Handler error: " + err.Error())
		}
		json.NewEncoder(w).Encode(respone)
		// w.WriteHeader(http.StatusOK)

	}

}
func CreateCategoryHandler(categorySvc category.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		category := dto.CategoryCreateRequest{}
		json.NewDecoder(r.Body).Decode(&category)

		response, err := categorySvc.CreateCategory(ctx, category)
		if err != nil {
			fmt.Println("error occured in createcategoryHandler" + err.Error())
		}
		json.NewEncoder(w).Encode(response)

	}

}
func UpdateCategoryHandler(categorySvc category.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		category := dto.Category{}
		json.NewDecoder(r.Body).Decode(&category)

		response, err := categorySvc.UpdateCategory(ctx, category)
		if err != nil {
			fmt.Println("error occured in createcategoryHandler" + err.Error())
		}
		json.NewEncoder(w).Encode(response)

	}

}
