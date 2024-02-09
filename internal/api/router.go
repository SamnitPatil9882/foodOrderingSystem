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

	//food
	r.HandleFunc("/foods", GetFoodListHandler(deps.FoodService)).Methods(http.MethodGet)
	r.HandleFunc("/foods/{category_id}", GetFoodListByCategoryHandler(deps.FoodService)).Methods(http.MethodGet)
	r.HandleFunc("/food", CreateFoodItemHandler(deps.FoodService)).Methods(http.MethodPost)
	r.HandleFunc("/food", UpdateFoodItemHandler(deps.FoodService)).Methods(http.MethodPatch)

	//user
	r.HandleFunc("/user/signup", SignUpHandler(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/user/login", LoginHandler(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/users", GetUsersHandler(deps.UserService)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", GetUserHandler(deps.UserService)).Methods(http.MethodGet)
	return r
}
