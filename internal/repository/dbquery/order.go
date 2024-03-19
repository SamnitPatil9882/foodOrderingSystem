package dbquery

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
	if role == "customer" {
		query = fmt.Sprintf("Select * FROM  `order` WHERE id = %d AND user_id = %d", orderID, userID)
	} else if role == "deliveryboy" {
		query = fmt.Sprintf(`SELECT o.id, o.user_id, o.created_at, o.total_amount, o.location
		FROM "order" o
		JOIN delivery d ON o.id = d.order_id
		WHERE d.deliveryboy_id = %d AND o.id = %d`, userID, orderID)
	} else {
		query = fmt.Sprintf("SELECT * FROM `order` WHERE id = %d", orderID)
	}
	rows, err := ordStr.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Printf("Error querying order by ID: %v", err)
		return ord, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ord.ID, &ord.UserID, &ord.CreatedAt, &ord.TotalAmout, &ord.Location)
		if err != nil {
			return dto.Order{}, fmt.Errorf("failed to scan order by ID: %w", err)
		}
	}
	return ord, nil
}

func (ordStr *OrderStore) GetOrderItemsByOrderID(ctx context.Context, orderID int, role string, userID int) ([]dto.CartItem, error) {
	var query string
	if role == "customer" {
		query = `SELECT oi.id AS orderitem_id, f.name AS foodname, f.price, oi.quantity
		FROM orderItem oi
		JOIN food f ON oi.food_id = f.id
		JOIN "order" o ON oi.order_id = o.id
		WHERE o.user_id = ? AND oi.order_id = ?;`
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
		return []dto.CartItem{}, err
	}
	defer rows.Close()

	var orderItems []dto.CartItem
	for rows.Next() {
		var oi dto.CartItem
		if err := rows.Scan(&oi.ID, &oi.FoodName, &oi.Price, &oi.Quantity); err != nil {
			log.Printf("Error scanning order item: %v", err)
			return []dto.CartItem{}, err
		}
		oi.Price = oi.Price * oi.Quantity
		orderItems = append(orderItems, oi)
	}
	return orderItems, nil
}
