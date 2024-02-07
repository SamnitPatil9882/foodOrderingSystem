package app

import (
	"database/sql"
	"net/http"

	"github.com/SamnitPatil9882/foodOrderingSystem/app/category"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, db *sql.DB) {

	//category
	r.HandleFunc("/menus", category.GetCategoriesHandler).Methods(http.MethodGet)
	r.HandleFunc("/menus/:menu_id", category.GetCategoryHandler).Methods(http.MethodGet)
	r.HandleFunc("/menus", category.CreateCategoryHandler).Methods(http.MethodPost)
	r.HandleFunc("/menus/:menu_id", category.UpdateCategoryHandler).Methods(http.MethodPatch)

	//
}
