package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/app/order"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func AddOrderItemHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req := dto.OrderItem{}
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "Request Invalid", http.StatusBadRequest)
			return
		}
		err = validateAddOrderItemReq(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = orderSvc.AddOrderItem(ctx, req.ID, req.Quantity)
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err)
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}
}
func GetOrderedItemsHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		resp := orderSvc.GetOrderList(ctx)
		if len(resp) == 0 {
			w.WriteHeader(http.StatusNotFound)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.EmptyCart}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func RemoveOrderedItemsHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id < 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidFoodID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		log.Println(id)
		resp, err := orderSvc.RemoveOrderItem(ctx, int(id))
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err)
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func CreateInvoiceHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		req := dto.InvoiceCreation{}
		err := json.NewDecoder(r.Body).Decode(&req)
		log.Println(req)
		if err != nil {
			http.Error(w, "enter valid data", http.StatusBadRequest)
			return
		}
		log.Println(req)
		err = validateCreateInvoice(&req)
		log.Print("validation result: ")
		// log.Println(err.Error())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("in validation if block")
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		userId := getUserID(r)
		log.Printf("user id : %d", userId)
		req.UserID = userId
		resp, err := orderSvc.CreateInvoice(ctx, req)
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err)
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetDeliveryListHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userID := getUserID(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid access token : order")
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		role := getRole(r)
		if role == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid access token: order role")
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if role == "customer" {
			resp, err := orderSvc.GetDeliveryList(ctx, userID)
			if err != nil {
				httpStatusCode := internal.GetHTTPStatusCode(err)
				w.WriteHeader(httpStatusCode)
				errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
				json.NewEncoder(w).Encode(errResp)
				return
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			resp, err := orderSvc.GetDeliveryList(ctx, -1)
			if err != nil {
				httpStatusCode := internal.GetHTTPStatusCode(err)
				w.WriteHeader(httpStatusCode)
				errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
				json.NewEncoder(w).Encode(errResp)
				return
			}
			json.NewEncoder(w).Encode(resp)
		}

	}
}

func GetDeliveryByIdHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidID}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		userID := getUserID(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid access token: order id")
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := orderSvc.GetDeliveryByID(ctx, userID, int(id))
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err);
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetListOfOrderHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userID := getUserID(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid access token: order id 1")
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		role := getRole(r)
		if role == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid access token: order role")
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		resp, err := orderSvc.GetListOfOrders(ctx, userID, role)
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err);
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		log.Printf("list of orders: %v", resp)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func GetOrderByIDHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		userID := getUserID(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid access token: order id")
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		role := getRole(r)
		if role == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid access token: order role")
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := orderSvc.GetOrderByID(ctx, int(id), userID, role)
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err);
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}

func GetInvoiceOrderByIDHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		userID := getUserID(r)
		if userID == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println("invalid access token: order id")
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		role := getRole(r)
		if role == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("invalid access token: order role")
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := orderSvc.GetInvoiceByOrderID(ctx, int(id), userID, role)
		if err != nil {
			httpStatusCode := internal.GetHTTPStatusCode(err);
			w.WriteHeader(httpStatusCode)
			errResp := dto.ErrorResponse{Error: httpStatusCode, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		json.NewEncoder(w).Encode(resp)
	}
}
func UpdateDeliveryHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		req := dto.DeliveryUpdateRequst{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = validateUpdateDelivery(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		err = orderSvc.UpdateDeliveryInfo(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
	}
}

func GetOrderItemsByOrderIDHandler(orderSvc order.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		params := mux.Vars(r)

		id, err := strconv.ParseInt(params["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: internal.InternalServerError}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		if id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			errResp := dto.ErrorResponse{Error: http.StatusBadRequest, Description: internal.InvalidID}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		userId := getUserID(r)
		role := getRole(r)
		if userId == 0 || role == "" {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.Unauthorized}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		resp, err := orderSvc.GetOrderItemsByOrderID(ctx, int(id), role, userId)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			errResp := dto.ErrorResponse{Error: http.StatusInternalServerError, Description: err.Error()}
			json.NewEncoder(w).Encode(errResp)
			return
		}
		log.Printf("orderitem by orderid : %v", resp)
		json.NewEncoder(w).Encode(resp)
	}
}
func getUserID(r *http.Request) int {
	log.Println("in get userid")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

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
func getRole(r *http.Request) string {
	log.Println("in get userid")
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse and validate JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWTKEY), nil
	})
	if err != nil {
		log.Println("Unauthorized: Invalid token claims")
		return ""
	}

	// Extract user roles from JWT claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Unauthorized: Invalid token claims")
		return ""
	}
	role := claims["role"].(string)
	return role
}
