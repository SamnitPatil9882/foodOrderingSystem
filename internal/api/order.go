package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/order"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AddOrderItemHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req := dto.OrderItem{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Request Invalid", http.StatusBadRequest)
			return
		}
		err = orderSvc.AddOrderItem(ctx, req.ID, req.Quantity)
		if err != nil {

			http.Error(w, "Request Invalid", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
func GetOrderedItemsHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resp := orderSvc.GetOrderList(ctx)
		if len(resp) == 0 {
			w.Write([]byte("empty cart"))
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func RemoveOrderedItemsHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println(id)
		resp, err := orderSvc.RemoveOrderItem(ctx, int(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func CreateInvoiceHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := dto.InvoiceCreation{}
		err := json.NewDecoder(r.Body).Decode(&req)
		fmt.Println(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		userId := getUserID(r)
		req.UserID = userId
		resp, err := orderSvc.CreateInvoice(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func getUserID(r *http.Request) int {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer", "", 1)

	// Parse and validate JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTKEY), nil
	})
	if err != nil {
		log.Println("Unauthorized: Invalid token claims")
		return 0
	}

	// Extract user roles from JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Unauthorized: Invalid token claims")
		return 0
	}
	userIdClaims := claims["userid"].(float64)
	userId := int(userIdClaims)
	return userId
}
