package api

import (
	"net/http"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {

	r := mux.NewRouter()
	//category
	r.HandleFunc("/categories", GetCategoriesHandler(deps.CategoryService)).Methods(http.MethodGet)
	r.HandleFunc("/categories/{category_id}", GetCategoryHandler(deps.CategoryService)).Methods(http.MethodGet)
	r.HandleFunc("/category", CreateCategoryHandler(deps.CategoryService)).Methods(http.MethodPost)
	r.HandleFunc("/category/{category_id}", UpdateCategoryHandler(deps.CategoryService)).Methods(http.MethodPatch)

	//d

	return r
}
