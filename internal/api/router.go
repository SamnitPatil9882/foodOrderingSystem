package api

import (
	"net/http"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(deps app.Dependencies) *mux.Router {


	
	
	r := mux.NewRouter()
	// handler:= c.Handler(mux)
	//category
	r.Handle("/categories", middleware.RequireAuth(GetCategoriesHandler(deps.CategoryService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/categories/{category_id}", middleware.RequireAuth(GetCategoryHandler(deps.CategoryService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/category", middleware.RequireAuth(CreateCategoryHandler(deps.CategoryService), []string{"admin"})).Methods(http.MethodPost)
	r.Handle("/category", middleware.RequireAuth(UpdateCategoryHandler(deps.CategoryService), []string{"admin"})).Methods(http.MethodPatch)

	//food
	r.Handle("/foods", middleware.RequireAuth(GetFoodListHandler(deps.FoodService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/foods/{category_id}", middleware.RequireAuth(GetFoodListByCategoryHandler(deps.FoodService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/food/{food_id}", middleware.RequireAuth(GetFoodInfoByIDHandler(deps.FoodService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/food", middleware.RequireAuth(CreateFoodItemHandler(deps.FoodService), []string{"admin"})).Methods(http.MethodPost)
	r.Handle("/food", middleware.RequireAuth(UpdateFoodItemHandler(deps.FoodService), []string{"admin"})).Methods(http.MethodPatch)

	//user
	r.HandleFunc("/user/signup", SignUpHandler(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/user/login", LoginHandler(deps.UserService)).Methods(http.MethodPost)
	r.Handle("/users", middleware.RequireAuth(GetUsersHandler(deps.UserService), []string{"admin"})).Methods(http.MethodGet)
	// r.HandleFunc("/users", GetUsersHandler(deps.UserService)).Methods(http.MethodGet)
	r.Handle("/user/{id}", middleware.RequireAuth(GetUserHandler(deps.UserService), []string{"admin"})).Methods(http.MethodGet)
	r.Handle("/user/update", middleware.RequireAuth(UpdateUserHandler(deps.UserService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodPut)

	// router.Handle("/admin/signup", middleware.RequireAuth(createUserHandler(deps.UserService), []string{"super_admin", "admin"})).Methods("POST")
	//orders
	// middleware.RequireAuth(GetOrderItemsByOrderIDHandler(deps.OrderService), []string{"admin"})
	r.Handle("/orderitem", middleware.RequireAuth(AddOrderItemHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodPost)
	r.Handle("/orderitem", middleware.RequireAuth(GetOrderedItemsHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/orderitem/remove/{id}", middleware.RequireAuth(RemoveOrderedItemsHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodDelete)
	r.Handle("/order/checkout", middleware.RequireAuth(CreateInvoiceHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodPost)
	r.Handle("/order/delivery", middleware.RequireAuth(UpdateDeliveryHandler(deps.OrderService), []string{"deliveryboy", "admin"})).Methods(http.MethodPut)

	r.Handle("/order/delivery/{id}", middleware.RequireAuth(GetDeliveryByIdHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/order/delivery", middleware.RequireAuth(GetDeliveryListHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)
	r.Handle("/orders", middleware.RequireAuth(GetListOfOrderHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)                  // c
	r.Handle("/orders/{id}", middleware.RequireAuth(GetOrderByIDHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)               // c
	r.Handle("/order/invoice/{id}", middleware.RequireAuth(GetInvoiceOrderByIDHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet) //
	r.Handle("/order/orderitem/{id}", middleware.RequireAuth(GetOrderItemsByOrderIDHandler(deps.OrderService), []string{"customer", "deliveryboy", "admin"})).Methods(http.MethodGet)

	// r.Use(middleware.JWTMiddleware)
	return r
}
