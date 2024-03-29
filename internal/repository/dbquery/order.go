package dbquery

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type OrderStore struct {
	BaseRepsitory
}

func NewOrderRepo(db *sql.DB) repository.OrderStorer {
	return &OrderStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (ordStr *OrderStore) GetListOfOrder(ctx context.Context, userID int, role string) ([]dto.Order, error) {
	fmt.Println("getlistoforder", userID, role);
	ordList := make([]dto.Order, 0)
	var query string
	if role == "customer" {
		query = fmt.Sprintf("SELECT * FROM `order` WHERE user_id = %d", userID)
	} else if role == "deliveryboy" {
		query = fmt.Sprintf(` SELECT o.id, o.user_id, o.created_at, o.total_amount, o.location
		FROM "order" AS o
		JOIN delivery AS d ON o.id = d.order_id
		WHERE d.deliveryboy_id = %d`, userID)
	} else {
		fmt.Printf("admin orders get")
		query = "SELECT * FROM `order` "
	}
	rows, err := ordStr.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Printf("Error querying orders: %v", err)
		return ordList, err
	}
	defer rows.Close()
	for rows.Next() {
		ord := dto.Order{}
		err := rows.Scan(&ord.ID, &ord.UserID, &ord.CreatedAt, &ord.TotalAmout, &ord.Location)
		if err != nil {
			return []dto.Order{}, fmt.Errorf("failed to scan order: %w", err)
		}
		ordList = append(ordList, ord)
	}
	return ordList, nil
}

func (ordStr *OrderStore) GetOrderByID(ctx context.Context, orderID int, userID int, role string) (dto.Order, error) {
	ord := dto.Order{}
	var query string

	orderExists, err := ordStr.orderExists(orderID, userID, role)
	if err != nil {
		return ord, err
	}
	if !orderExists {
		return ord, internal.ErrOrderIdNotExists;
	}
	if role == "customer" {
		query = fmt.Sprintf("SELECT * FROM `order` WHERE id = %d AND user_id = %d", orderID, userID)
	} else if role == "deliveryboy" {
		query = fmt.Sprintf(`SELECT o.id, o.user_id, o.created_at, o.total_amount, o.location
		FROM "order" o
		JOIN delivery d ON o.id = d.order_id
		WHERE d.deliveryboy_id = %d AND o.id = %d`, userID, orderID)
	} else {
		query = fmt.Sprintf("SELECT * FROM `order` WHERE id = %d", orderID)
	}
	rows, err := ordStr.BaseRepsitory.DB.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying order by ID: %v", err)
		return ord, err
	}
	defer rows.Close()

	// Check if any rows were returned
	if !rows.Next() {
		return ord, internal.ErrOrderNotFound
	}

	// Scan the result
	err = rows.Scan(&ord.ID, &ord.UserID, &ord.CreatedAt, &ord.TotalAmout, &ord.Location)
	if err != nil {
		log.Printf("Failed to scan order by ID: %v", err)
		return ord, fmt.Errorf("failed to scan order by ID: %w", err)
	}

	return ord, nil
}

func (ordStr *OrderStore) GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error) {
	// Check if the order with the specified ID exists
	orderExists, err := ordStr.orderExists(orderID, userID, role)
	if err != nil {
		return nil, err
	}
	if !orderExists {
		return nil, internal.ErrOrderIdNotExists;
	}

	var query string
	if role == "customer" {

		fmt.Println("customer: orderitem: ",orderID," ",role," ",userID)
		query = `SELECT oi.id AS orderitem_id, f.name AS foodname, f.price, oi.quantity
		FROM orderItem oi
		JOIN food f ON oi.food_id = f.id
		JOIN "order" o ON oi.order_id = o.id
		WHERE  oi.order_id = ? AND o.user_id = ?;`
		// query = `SELECT oi.id AS orderitem_id,  f.name AS foodname, f.price,oi.quantity
        //   FROM orderItem oi
        //   JOIN food f ON oi.food_id = f.id
        //   JOIN "order" o ON oi.order_id = o.id
        //   WHERE o.id = ? AND o.user_id = ?`
		
	} else if role == "deliveryboy" {
		query = `
		SELECT oi.id AS orderitem_id, f.name AS food_name, f.price AS food_price, oi.quantity
		FROM orderItem oi
		JOIN food f ON oi.food_id = f.id
		JOIN "order" o ON oi.order_id = o.id
		JOIN delivery d ON d.order_id = o.id
		WHERE d.deliveryboy_id = ? AND oi.order_id = ?;
		`
	} else {
		query = `
		SELECT oi.id AS orderitem_id, f.name AS food_name, f.price, oi.quantity
		FROM orderItem oi
		JOIN food f ON oi.food_id = f.id
		WHERE oi.order_id = ?;
		`
	}

	rows, err := ordStr.BaseRepsitory.DB.Query(query, orderID, userID)
	if err != nil {
		log.Printf("Error querying order items: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orderItems []dto.CartItem
	for rows.Next() {
		var oi dto.CartItem
		if err := rows.Scan(&oi.ID, &oi.FoodName, &oi.Price, &oi.Quantity); err != nil {
			log.Printf("Error scanning order item: %v", err)
			return nil, err
		}
		oi.Price = oi.Price * oi.Quantity
		orderItems = append(orderItems, oi)
	}
	fmt.Println("order items: ",orderItems)
	return orderItems, nil
}

// Helper function to check if the order exists
func (ordStr *OrderStore) orderExists(orderID int, userID int, role string) (bool, error) {
	var count int
	var query string
	if role == "customer" {
		query = "SELECT COUNT(*) FROM `order` WHERE id = ? AND user_id = ?"
	} else if role == "deliveryboy" {
		query = `SELECT COUNT(*)
		FROM "order" o
		JOIN delivery d ON o.id = d.order_id
		WHERE d.deliveryboy_id = ? AND o.id = ?`
	} else {
		query = "SELECT COUNT(*) FROM `order` WHERE id = ?"
	}

	err := ordStr.BaseRepsitory.DB.QueryRow(query, orderID, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
