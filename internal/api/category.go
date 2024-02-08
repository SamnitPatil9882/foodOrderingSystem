package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/category"
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
func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {

}
func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {

}
