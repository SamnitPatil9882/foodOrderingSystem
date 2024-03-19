package main

import (
	"fmt"
	"net/http"

	// "github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/api"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

func main() {

	//context creation
	// ctx := context.Background()

	//database creation
	database, err := repository.InitializeDatabase()
	if err != nil {
		fmt.Println("error occured in creation of db")
	}

	//intialize service dependencies
	services := app.NewServices(database)

	//intialize router
	r := api.NewRouter(services)

	http.ListenAndServe(fmt.Sprintf(":%d", constants.HTTPPort), r)
	fmt.Println("Server is running on port : ", constants.HTTPPort)

}
