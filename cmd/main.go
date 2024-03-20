package main

import (
	"fmt"
	"net/http"

	// "github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/api"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
	"github.com/rs/cors"
)

func main() {

	//context creation
	// ctx := context.Background()

	//database creation
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
	})
	database, err := repository.InitializeDatabase()
	if err != nil {
		fmt.Println("error occured in creation of db")
	}

	
	//intialize service dependencies
	services := app.NewServices(database)
	r := api.NewRouter(services)
	handler := c.Handler(r)

	fmt.Println("Server is running on port : ", constants.HTTPPort)
	http.ListenAndServe(fmt.Sprintf(":%d", constants.HTTPPort), handler)

}
