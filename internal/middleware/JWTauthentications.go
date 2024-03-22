package middleware

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/constants"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
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

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

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
		return role == "deliveryboy" || role == "admin"

	default:
		return false
	}
}

// /*****
func RequireAuth(next http.Handler, roles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse and validate JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user roles from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Create User object with extracted roles
		// user := User{
		// 	Username: claims["username"].(string),
		// 	Roles:    claims["role"],
		// }
		// userId:=claims["userid"].(float64)

		role := claims["role"].(string)

		if !slices.Contains(roles, role) {
			w.WriteHeader(http.StatusUnauthorized)
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: internal.UnauthorizedAccess}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		// r.Header.Set("user_id", Id.String())
		// r.Header.Set("role", Role)

		next.ServeHTTP(w, r)

	})

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

/*
type AccessRule struct {
	Role   string
	Method string
	Path   string
}

var accessRules = []AccessRule{
	{Role: "admin", Method: "GET", Path: "/users"},
	{Role: "admin", Method: "POST", Path: "/users"},
	{Role: "deliveryboy", Method: "GET", Path: "/user/{id}"},
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		// Parse and validate JWT token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user role from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}
		role := claims["role"].(string)

		// Get request method and path
		method := r.Method
		path := r.URL.Path

		// Check if the user's role is allowed for the requested URL and method
		authorized := false
		for _, rule := range accessRules {
			if rule.Role == role && rule.Method == method && matchPath(rule.Path, path) {
				authorized = true
				break
			}
		}

		if !authorized {
			http.Error(w, "Forbidden: Insufficient privileges", http.StatusForbidden)
			return
		}

		// User is authorized, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

// matchPath checks if the given URL path matches the pattern, which may contain wildcards.
func matchPath(pattern, path string) bool {
	if pattern == path {
		return true
	}

	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i, part := range patternParts {
		if part != pathParts[i] && part != "{id}" {
			return false
		}
	}

	return true
}


*/
/*type AccessRule struct {
	Role   string
	Method string
	Path   string
}

var accessRules = []AccessRule{
	{Role: "admin", Method: "GET", Path: "/users"},
	{Role: "admin", Method: "POST", Path: "/users"},
	{Role: "deliveryboy", Method: "GET", Path: "/user/{id}"},
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract JWT token from Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized: Token missing", http.StatusUnauthorized)
			return
		}

		// Parse and validate JWT token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.JWTKEY), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user role from JWT claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized: Invalid token claims", http.StatusUnauthorized)
			return
		}
		role := claims["role"].(string)

		// Get request method and path
		method := r.Method
		path := r.URL.Path

		// Check if the user's role is allowed for the requested URL and method
		authorized := false
		for _, rule := range accessRules {
			if rule.Role == role && rule.Method == method && matchPath(rule.Path, path) {
				authorized = true
				break
			}
		}

		if !authorized {
			errResp := dto.ErrorResponse{Error: http.StatusUnauthorized, Description: "Enter Valid Credentials"}
			json.NewEncoder(w).Encode(errResp)
			return
		}

		// User is authorized, pass control to the next handler
		next.ServeHTTP(w, r)
	})
}

// matchPath checks if the given URL path matches the pattern, which may contain wildcards.
func matchPath(pattern, path string) bool {
	if pattern == path {
		return true
	}

	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i, part := range patternParts {
		if part != pathParts[i] && part != "{id}" {
			return false
		}
	}

	return true
}*/
