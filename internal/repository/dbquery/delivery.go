package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/repository"
)

type DeliveryStore struct {
	BaseRepsitory
}

func NewDeliveryRepo(db *sql.DB) repository.DeliveryStorer {
	return &DeliveryStore{
		BaseRepsitory: BaseRepsitory{db},
	}
}

func (dlvstr *DeliveryStore) GetDeliveryList(ctx context.Context, userID int) ([]dto.Delivery, error) {
	dlvList := make([]dto.Delivery, 0)
	var query string
	if userID != -1 {
		query = fmt.Sprintf(`
		SELECT *
		FROM delivery d
		JOIN "order" o ON d.order_id = o.id
		WHERE o.user_id = %d
		`, userID)

	} else {
		query = "SELECT * FROM delivery "
	}

	rows, err := dlvstr.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println(err)
		return dlvList, err
	}
	defer rows.Close()
	for rows.Next() {
		dlvry := dto.Delivery{}
		var deliveryBoyID sql.NullInt64
		var endAt sql.NullString
		err := rows.Scan(&dlvry.ID, &dlvry.OrderID, &deliveryBoyID, &dlvry.StartTime, &endAt, &dlvry.Status)
		if err != nil {
			log.Println(err)
			return dlvList, err
		}
		if deliveryBoyID.Valid {
			dlvry.UserID = int(deliveryBoyID.Int64)
		} else {
			dlvry.UserID = 0
		}
		if endAt.Valid {
			dlvry.EndTime = endAt.String
		} else {
			dlvry.EndTime = ""
		}

		dlvList = append(dlvList, dlvry)
	}
	return dlvList, nil
}

func (dlvstr *DeliveryStore) UpdateDeliveryInfo(ctx context.Context, updateInfo dto.DeliveryUpdateRequst) error {
	query := fmt.Sprintf("SELECT * FROM user WHERE id = %d AND role = '%s'", updateInfo.UserID, "deliveryboy")

	rows, err := dlvstr.BaseRepsitory.DB.Query(query)
	if err != nil {
		log.Println(err)
		// return errors.New("failed to query user information")
		return err;
	}
	defer rows.Close()
	flag := false
	for rows.Next() {
		flag = true
		rows.Close()
		break
	}
	if !flag {
		return internal.ErrDeliveryBoyIdNotExists
	}

	// Fetch the current status of the delivery
	currentStatus, err := dlvstr.getCurrentStatus(updateInfo.ID)
	if err != nil {
		return internal.ErrInvalidDeliveryId
	}

	// Check if the delivery status is transitioning to "pickup"
	if updateInfo.Status == "pickup" {
		// If the current status is not "preparing", return an error
		if currentStatus != "preparing" {
			// return errors.New("cannot transition to 'pickup' from current status")
			return internal.ErrInvalidDeliveryStatusToPickup
		}
	}

	// Check if the delivery status is transitioning to "delivered"
	if updateInfo.Status == "delivered" {
		// If the current status is not "pickup", return an error
		if currentStatus != "pickup" {
			// return errors.New("cannot transition to 'delivered' from current status")
			return internal.ErrInvalidDeliveryStatusToDelivered
		}

		// Update end time to current time
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		query = "UPDATE delivery SET deliveryboy_id=?, end_at=?, status=? WHERE id=?"
		statement, err := dlvstr.BaseRepsitory.DB.Prepare(query)
		if err != nil {
			log.Println("error occurred in updating delivery db: " + err.Error())
			return errors.New("failed to prepare delivery update statement"+err.Error())
		}
		defer statement.Close()

		res, err := statement.Exec(updateInfo.UserID, currentTime, updateInfo.Status, updateInfo.ID)
		if err != nil {
			log.Println("error occurred in updating delivery db: " + err.Error())
			return errors.New("failed to execute delivery update statement")
		}
		noOfRawAffected, err := res.RowsAffected()
		if err != nil {
			return errors.New("failed to get affected rows after delivery update")
		}
		if noOfRawAffected == 0 {
			return errors.New("no rows affected after delivery update")
		}
	}

	// Update deliveryboy_id and status for "pickup"
	if updateInfo.Status == "pickup" {
		query = "UPDATE delivery SET deliveryboy_id=?, status=? WHERE id=?"
		statement, err := dlvstr.BaseRepsitory.DB.Prepare(query)
		if err != nil {
			log.Println("error occurred in updating delivery db: " + err.Error())
			return errors.New("failed to prepare delivery update statement")
		}
		defer statement.Close()

		res, err := statement.Exec(updateInfo.UserID, updateInfo.Status, updateInfo.ID)
		if err != nil {
			log.Println("error occurred in updating delivery db: " + err.Error())
			return errors.New("failed to execute delivery update statement")
		}
		noOfRawAffected, err := res.RowsAffected()
		if err != nil {
			return errors.New("failed to get affected rows after delivery update")
		}
		if noOfRawAffected == 0 {
			return errors.New("no rows affected after delivery update")
		}
	}

	return nil
}

func (dlvstr *DeliveryStore) getCurrentStatus(deliveryID int) (string, error) {
	var currentStatus string
	query := "SELECT status FROM delivery WHERE id = ?"
	err := dlvstr.BaseRepsitory.DB.QueryRow(query, deliveryID).Scan(&currentStatus)
	if err != nil {
		log.Println(err)
		return "", errors.New("failed to fetch current delivery status")
	}
	return currentStatus, nil
}

func (dlvstr *DeliveryStore) GetDeliveryByID(ctx context.Context, userID int, deliveryID int) (dto.Delivery, error) {
	var delivery dto.Delivery

	// Check the role of the user
	userRole, err := dlvstr.getUserRole(userID)
	if err != nil {
		return delivery, errors.New("failed to fetch user role")
	}

	if !dlvstr.deliveryIDExists(deliveryID) {
		return delivery, internal.ErrOrderIdNotExists
	}

	switch userRole {
	case "admin":
		// Admin gets all delivery info by ID
		query := "SELECT * FROM delivery WHERE id = ?"
		err := dlvstr.BaseRepsitory.DB.QueryRow(query, deliveryID).Scan(&delivery.ID, &delivery.OrderID, &delivery.UserID, &delivery.StartTime, &delivery.EndTime, &delivery.Status)
		if err != nil {
			log.Println(err)
			return delivery, errors.New("failed to fetch delivery info for admin")
		}
	case "customer":
		// Customers get delivery info for their own orders
		query := `
            SELECT d.id, d.order_id, d.deliveryboy_id, d.start_at, d.end_at, d.status
            FROM delivery d
            JOIN "order" o ON d.order_id = o.id
            WHERE o.user_id = ? AND d.id = ?
        `
		err := dlvstr.BaseRepsitory.DB.QueryRow(query, userID, deliveryID).Scan(&delivery.ID, &delivery.OrderID, &delivery.UserID, &delivery.StartTime, &delivery.EndTime, &delivery.Status)
		if err != nil {
			log.Println(err)
			return delivery, errors.New("failed to fetch delivery info for customer")
		}
	case "deliveryboy":
		// Delivery boys get delivery info for their assigned orders
		query := `
            SELECT d.id, d.order_id, d.deliveryboy_id, d.start_at, d.end_at, d.status
            FROM delivery d
            WHERE d.deliveryboy_id = ? AND d.id = ?
        `
		err := dlvstr.BaseRepsitory.DB.QueryRow(query, userID, deliveryID).Scan(&delivery.ID, &delivery.OrderID, &delivery.UserID, &delivery.StartTime, &delivery.EndTime, &delivery.Status)
		if err != nil {
			log.Println(err)
			return delivery, errors.New("failed to fetch delivery info for delivery boy")
		}
	default:
		return delivery, errors.New("invalid user role")
	}

	return delivery, nil
}

func (dlvstr *DeliveryStore) deliveryIDExists(deliveryID int) bool {
	query := "SELECT COUNT(*) FROM delivery WHERE id = ?"
	var count int
	err := dlvstr.BaseRepsitory.DB.QueryRow(query, deliveryID).Scan(&count)
	if err != nil {
		log.Println("Error checking if delivery ID exists:", err)
		return false
	}
	return count > 0
}
// Helper function to get user role
func (dlvstr *DeliveryStore) getUserRole(userID int) (string, error) {
	var userRole string
	query := "SELECT role FROM user WHERE id = ?"
	err := dlvstr.BaseRepsitory.DB.QueryRow(query, userID).Scan(&userRole)
	if err != nil {
		log.Println(err)
		return "", errors.New("failed to fetch user role")
	}
	return userRole, nil
}
