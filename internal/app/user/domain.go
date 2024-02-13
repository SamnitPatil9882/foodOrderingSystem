package user

import (
	"fmt"
	"regexp"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func validateUser(user *dto.UserSignUpRequest) bool {
	if len(user.Firstname) < 2 {
		return false
	} else if len(user.Lastname) < 2 {
		return false
	} else if len(user.Phone) < 10 {
		return false
	} else if user.Role != "deliveryboy" && user.Role != "customer" && user.Role != "admin" {
		return false
	} else if !(isValidEmail(user.Email) && isValidPassword(user.Password)) {
		return false
	}
	user.Lastname = cases.Title(language.Und, cases.NoLower).String(user.Lastname)
	user.Firstname = cases.Title(language.Und, cases.NoLower).String(user.Firstname)
	return true
}

func isValidEmail(email string) bool {
	// Regular expression for validating email addresses
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		fmt.Println("invalid email")
	}
	return re.MatchString(email)
}
func isValidPassword(password string) bool {
	if len(password) < 10 {
		return false
	}
	passwordRegex := `^(.*[a-zA-Z].*[a-zA-Z].*[a-zA-Z].*[a-zA-Z].*[a-zA-Z])(.*[!@#$%^&*()-_+=?])(.*[0-9].*[0-9].*[0-9])`
	re := regexp.MustCompile(passwordRegex)
	if !re.MatchString(password) {
		fmt.Println("invalid password")
	}
	return re.MatchString(password)
}
