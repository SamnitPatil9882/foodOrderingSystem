package dbquery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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
		return err
	}
	defer rows.Close()
	flag := false
	for rows.Next() {
		flag = true
		rows.Close()
		break
	}
	if !flag {
		return errors.New("enter user id of delivery boy")
	}

	// Check if the delivery status is transitioning to "pickup"
	if updateInfo.Status == "pickup" {
		// Fetch the current status of the delivery
		currentStatus, err := dlvstr.getCurrentStatus(updateInfo.ID)
		if err != nil {
			return err
		}

		// If the current status is not "preparing", return an error
		if currentStatus != "preparing" {
			return errors.New("cannot transition to 'pickup' from current status")
		}
	}

	// Check if the delivery status is transitioning to "delivered"
	if updateInfo.Status == "delivered" {
		// Fetch the current status of the delivery
		currentStatus, err := dlvstr.getCurrentStatus(updateInfo.ID)
		if err != nil {
			return err
		}

		// If the current status is not "pickup", return an error
		if currentStatus != "pickup" {
			return errors.New("cannot transition to 'delivered' from current status")
		}
	}
	query = "UPDATE delivery SET deliveryboy_id=?, end_at=?, status=? WHERE id=?"

	statement, err := dlvstr.BaseRepsitory.DB.Prepare(query)
	if err != nil {
		log.Println("error occurred in updating delivery db: " + err.Error())
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(updateInfo.UserID, updateInfo.EndTime, updateInfo.Status, updateInfo.ID)
	if err != nil {
		log.Println("error occurred in updating delivery db: " + err.Error())
		return err
	}
	return nil
	/*
		query := fmt.Sprintf("SELECT * FROM user WHERE id = %d", updateInfo.UserID)
		rows, err := dlvstr.BaseRepsitory.DB.Query(query)
		if err != nil {
			log.Println(err)
			return err
		}
		defer rows.Close()
		flag := false
		for rows.Next() {
			flag = true
			break
		}
		if !flag {
			return errors.New("enter user id of delivery boy")
		}
		if len(updateInfo.EndTime) != 0 {
			query := fmt.Sprintf("SELECT start_at FROM delivery WHERE id = %d", updateInfo.ID)
			rows, err := dlvstr.BaseRepsitory.DB.Query(query)
			if err != nil {
				log.Println(err)
				return err
			}
			defer rows.Close()
			var startAt string
			for rows.Next() {
				err := rows.Scan(&startAt)
				if err != nil {
					log.Println(err)
					break
				}
			}

			endAt := updateInfo.EndTime
			log.Printf("start at: %s  ; end at : %s", startAt, endAt)
			layout := "2006-01-02 15:04:05"
			starttime, err := time.Parse(layout, startAt)
			if err != nil {
				log.Println("Error parsing start time:", err)
			}
			endtime, err := time.Parse(layout, endAt)
			if err != nil {
				log.Println("Error parsing end time:", err)
				return errors.New("enter valid end time")
			}

			if endtime.After(starttime) {
				log.Println("End time is greater than start time")
			} else if endtime.Before(starttime) {
				log.Println("End time is less than start time")
				return errors.New("end time is less than start time")
			} else {
				fmt.Println("End time is equal to start time")
				return errors.New("end time is equal to start time")
			}
		}

		query = "UPDATE delivery SET deliveryboy_id=?, end_at=?, status=? WHERE id=?"
		stmt, err := dlvstr.BaseRepsitory.DB.Prepare(query)
		if err != nil {
			log.Println("error occured in prepareing updating delivery db: " + err.Error())
			return err
		}
		defer stmt.Close()

		// Assuming updateInfo.UserID, updateInfo.EndTime, updateInfo.Status, and updateInfo.ID are the values you want to update
		_, err = stmt.Exec(updateInfo.UserID, updateInfo.EndTime, updateInfo.Status, updateInfo.ID)
		if err != nil {
			log.Println("error occured in updating delivery db: " + err.Error())
			return err
		}
		return nil*/
}

func (dlvstr *DeliveryStore) getCurrentStatus(deliveryID int) (string, error) {
	var currentStatus string
	query := "SELECT status FROM delivery WHERE id = ?"
	err := dlvstr.BaseRepsitory.DB.QueryRow(query, deliveryID).Scan(&currentStatus)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return currentStatus, nil
}
