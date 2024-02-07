package main

import (
	"fmt"

	"github.com/SamnitPatil9882/foodOrderingSystem/repository"
)

func main() {

	//context creation
	// ctx := context.Background()

	//database creation
	database, err := repository.InitializeDatabase()
	if err != nil {
		fmt.Println("error occured in creation of db")
	}

	// r := mux.NewRouter()

}
