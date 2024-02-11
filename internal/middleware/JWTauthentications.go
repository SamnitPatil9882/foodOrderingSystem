package middleware

import (
	"net/http"
	"strings"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/dgrijalva/jwt-go"
)

// Role constants
const (
	RoleCustomer    = "customer"
	RoleDeliveryBoy = "deliveryboy"
	RoleAdmin       = "admin"
)

// User struct to represent user information
type User struct {
	Username string
	Roles    []string
}

// JWT middleware to verify and extract user roles
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from Authorization header
		path := r.URL.Path
		if path == "/user/signup" || path == "/user/login" {
			next.ServeHTTP(w, r)
			return
		}
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer", "", 1)

		// Parse and validate JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user roles from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Create User object with extracted roles
		// user := User{
		// 	Username: claims["username"].(string),
		// 	Roles:    claims["role"],
		// }
		role := claims["role"].(string)

		// Check access based on user roles
		if role == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		if !hasAccess(role, path) {
			http.Error(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
			return
		}

		// Pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

func hasAccess(role string, path string) bool {
	switch path {
	case "/users":
		return role == "admin"

	case "/user/{id}":
		return role == "deliveryboy"
	default:
		return false
	}
}

// func getRole

// Helper function to check if a string slice contains a specific value
// func contains(role string, value string) bool {
// 	if role == value {
// 		return true
// 	}
// 	return false
// }

//	r.Use(jwtMiddleware)
