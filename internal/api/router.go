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
	// r.HandleFunc("/menus/:menu_id", category.GetCategoryHandler).Methods(http.MethodGet)
	// r.HandleFunc("/menus", category.CreateCategoryHandler).Methods(http.MethodPost)
	// r.HandleFunc("/menus/:menu_id", category.UpdateCategoryHandler).Methods(http.MethodPatch)

	//

	return r
}
