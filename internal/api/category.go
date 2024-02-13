package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
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
			log.Println("Handler error: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
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

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// http.Error(w, "enter valid data", http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if category_id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			// http.Error(w, "enter valid data", http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidCategoryID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		respone, err := categorySvc.GetCategory(ctx, category_id)
		if err != nil {
			log.Println("error in getting categories: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
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
		err := validateCreateCategoryReq(&category)
		if err != nil {
			log.Println("error in create category request")
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		response, err := categorySvc.CreateCategory(ctx, category)
		if err != nil {
			log.Println("error occured in createcategoryHandler" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(response)

	}

}
func UpdateCategoryHandler(categorySvc category.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		category := dto.Category{}
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			log.Println("error in request")
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = validateUpdateCategoryReq(&category)
		if err != nil {
			log.Println("error in create category request")
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		response, err := categorySvc.UpdateCategory(ctx, category)
		if err != nil {
			log.Println("error occured in createcategoryHandler: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(response)

	}

}

// func UpdateCategoryStatusHandler(categorySvc category.Service)func(w http.ResponseWriter,r *http.Request){
// 	return func(w http.ResponseWriter, r *http.Request){
// 		ctx := r.Context()
// 		category := dto.Category{}

// 		response,err :=categorySvc.UpdateCategoryStatus(ctx,id,)
// 	}
// }
