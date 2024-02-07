package main

import (
	"fmt"

	"github.com/SamnitPatil9882/foodOrderingSystem/app"
	"github.com/SamnitPatil9882/foodOrderingSystem/repository"
	"github.com/gorilla/mux"
)

func main() {

	//context creation
	// ctx := context.Background()

	//database creation
	database, err := repository.InitializeDatabase()
	if err != nil {
		fmt.Println("error occured in creation of db")
	}

	r := mux.NewRouter()
	app.Routes(r, database)

}
