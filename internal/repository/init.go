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
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println("Error enabling foreign key constraints:", err)
		return db, err
	}

	query := "CREATE TABLE IF NOT EXISTS category  (id INTEGER PRIMARY KEY AUTOINCREMENT,name VARCHAR(100) NOT NULL UNIQUE,description VARCHAR(255),is_active INTEGER DEFAULT 1 NOT NULL)"
	statement, err := database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of menu table: " + err.Error())
		return database, err
	}
	statement.Exec()
	seedCategoryData()
	query = "CREATE TABLE IF NOT EXISTS food  (id INTEGER PRIMARY KEY AUTOINCREMENT,category_id VARCHAR(100) NOT NULL,price INTEGER NOT NULL,name VARCHAR(100) NOT NULL UNIQUE,is_veg INTEGER NOT NULL DEFAULT 1,is_avail INTEGER NOT NULL DEFAULT 1,FOREIGN KEY (category_id) REFERENCES category(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of food table: " + err.Error())
		return database, err
	}
	statement.Exec()
	seedFoodData()
	query = "CREATE TABLE IF NOT EXISTS user  (id INTEGER PRIMARY KEY AUTOINCREMENT,phone VARCHAR(15) UNIQUE NOT NULL,email VARCHAR(255) UNIQUE NOT NULL,password VARCHAR(1000) NOT NULL,firstname VARCHAR(100) NOT NULL,lastname VARCHAR(100) NOT NULL,role VARCHAR(20) DEFAULT 'customer' CHECK(role IN ('customer','admin','deliveryboy')))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of user table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS \"order\" (id INTEGER PRIMARY KEY AUTOINCREMENT,user_id INTEGER NOT NULL,created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,total_amount INTEGER NOT NULL,location VARCHAR(500),FOREIGN KEY (user_id) REFERENCES user(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of order table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS orderItem  (id INTEGER NOT NULL ,order_id INTEGER NOT NULL,food_id INTEGER NOT NULL,quantity INTEGER NOT NULL,PRIMARY KEY (id,order_id),FOREIGN KEY (order_id) REFERENCES \"order\"(id),FOREIGN KEY (food_id) REFERENCES food(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of order item table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS invoice (id INTEGER PRIMARY KEY AUTOINCREMENT,order_id INTEGER NOT NULL,payment_method VARCHAR(100) NOT NULL DEFAULT 'creditcard' CHECK(payment_method IN ('creditcard','debitcard','visa','upi')),created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,FOREIGN KEY (order_id) REFERENCES \"order\"(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of  invoice table: " + err.Error())
		return database, err
	}
	statement.Exec()

	query = "CREATE TABLE IF NOT EXISTS delivery (id INTEGER PRIMARY KEY AUTOINCREMENT,order_id INTEGER NOT NULL,deliveryboy_id INTEGER ,start_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,end_at TIMESTAMP ,status VARCHAR(50) NOT NULL DEFAULT 'preaparing' CHECK(status IN ('preparing','pickup','delivered')),FOREIGN KEY (deliveryboy_id) REFERENCES user(id),FOREIGN KEY (order_id) REFERENCES `order`(id))"
	statement, err = database.Prepare(query)
	if err != nil {
		fmt.Println("error occured in creation of delivery table: " + err.Error())
		return database, err
	}
	statement.Exec()
	return database, nil
}

func seedCategoryData() {
	query := "INSERT INTO category (name,description,is_active) VALUES(?,?,?)"
	statement, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return
	}
	defer statement.Close()
	statement.Exec("maincourse", "abc", 1)
	statement.Exec("starter", "abc", 0)
	statement.Exec("softdrinks", "abc", 0)
	statement.Exec("dessert", "abc", 1)

}
func seedFoodData() {
	query := "INSERT INTO food (name,price,category_id,is_veg,is_avail) VALUES(?,?,?,?,?)"
	statement, err := db.Prepare(query)
	if err != nil {
		fmt.Println("error in inserting: " + err.Error())
		return
	}
	defer statement.Close()
	statement.Exec("roti", 25, 1, 1, 1)
	statement.Exec("panner", 100, 4, 1, 1)
	statement.Exec("Biryani", 150, 2, 0, 0)
	statement.Exec("orange juice", 50, 4, 1, 1)

}
