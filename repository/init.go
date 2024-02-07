package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDatabase() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "../foodOrderingDB.db")
	if err != nil {
		fmt.Println("error occured in creation of db" + err.Error())
		return database, err
	}
	db = database

	query := "CREATE TABLE IF NOT EXISTS category  (id INTEGER PRIMARY KEY AUTOINCREMENT,category_name VARCHAR(100) NOT NULL,description VARCHAR(255),is_active INTEGER DEFAULT 1)"
	statement, err := database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of menu table: " + err.Error())
		return database, err
	}
	statement.Exec()
	seedCategoryData()
	query = "CREATE TABLE IF NOT EXISTS food  (id INTEGER PRIMARY KEY AUTOINCREMENT,category_id VARCHAR(100) NOT NULL,price INTEGER,name VARCHAR(100),is_veg INTEGER DEFAULT 1,FOREIGN KEY (category_id) REFERENCES category(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of food table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS user  (id INTEGER PRIMARY KEY AUTOINCREMENT,phone VARCHAR(15) UNIQUE NOT NULL,email VARCHAR(255) UNIQUE NOT NULL,password VARCHAR(255) NOT NULL,firstname VARCHAR(100),lastname VARCHAR(100),role VARCHAR(20) DEFAULT 'customer' CHECK(role IN ('customer','admin')))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of user table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS \"order\" (id INTEGER PRIMARY KEY AUTOINCREMENT,user_id INTEGER NOT NULL,order_date DATE NOT NULL,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,total_amount INTEGER NOT NULL,FOREIGN KEY (user_id) REFERENCES user(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of order table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS orderItem  (id INTEGER PRIMARY KEY AUTOINCREMENT,order_id INTEGER NOT NULL,food_id INTEGER NOT NULL,quantity INTEGER NOT NULL,FOREIGN KEY (order_id) REFERENCES \"order\"(id),FOREIGN KEY (food_id) REFERENCES food(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of order item table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS invoice (id INTEGER PRIMARY KEY AUTOINCREMENT,order_id INTEGER NOT NULL,total_amount INTEGER NOT NULL,date DATE NOT NULL,payment_method VARCHAR(100) NOT NULL,FOREIGN KEY (order_id) REFERENCES \"order\"(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of  invoice table: " + err.Error())
		return database, err
	}
	statement.Exec()

	return database, nil
}

func seedCategoryData() {
	query := "INSERT INTO category (category_name,description,is_active) VALUES(?,?,?)"
	statement, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return
	}
	statement.Exec("maincourse", "abc", 1)
	statement.Exec("starter", "abc", 0)
	statement.Exec("softdrinks", "abc", 0)
	statement.Exec("dessert", "abc", 1)
}
