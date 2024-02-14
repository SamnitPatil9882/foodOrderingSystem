package api

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/SamnitPatil9882/foodOrderingSystem/internal"
	"github.com/SamnitPatil9882/foodOrderingSystem/internal/pkg/dto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var paymentMethodTypes = map[string]int{
	"creditcard": 1,
	"debitcard":  1,
	"upi":        1,
	"visa":       1,
}
var deliveryStatus = map[string]int{
	"pickup":    1,
	"preparing": 1,
	"delivered": 1,
}
var roleOfUser = map[string]int{
	"customer":    1,
	"deliveryboy": 1,
	"admin":       1,
}

func validateCreateCategoryReq(createCategory *dto.CategoryCreateRequest) error {
	if len(createCategory.Name) < 2 {
		return errors.New(internal.InvalidCategoryName)
	}
	createCategory.Name = cases.Title(language.Und, cases.NoLower).String(createCategory.Name)
	createCategory.Description = cases.Title(language.Und, cases.NoLower).String(createCategory.Description)
	if createCategory.IsActive < 0 {
		createCategory.IsActive = 0
	} else if createCategory.IsActive > 1 {
		createCategory.IsActive = 1
	}
	return nil
}
func validateUpdateCategoryReq(updateCategory *dto.Category) error {
	if updateCategory.ID <= 0 {
		return errors.New(internal.InvalidCategoryID)
	} else if len(updateCategory.Name) < 2 {
		return errors.New(internal.InvalidCategoryName)
	}
	updateCategory.Name = cases.Title(language.Und, cases.NoLower).String(updateCategory.Name)
	updateCategory.Description = cases.Title(language.Und, cases.NoLower).String(updateCategory.Description)
	if updateCategory.IsActive < 0 {
		updateCategory.IsActive = 0
	} else if updateCategory.IsActive > 1 {
		updateCategory.IsActive = 1
	} else if updateCategory.IsActive != 0 && updateCategory.IsActive != 1 {
		updateCategory.IsActive = 1
	}
	return nil
}

func validateFoodCreateReq(createFood *dto.FoodCreateRequest) error {
	if createFood.CategoryID <= 0 || createFood.Price <= 0 {
		return errors.New(internal.InvalidCategoryID)
	} else if len(createFood.Name) <= 2 {
		return errors.New(internal.InvalidFoodName)
	}

	createFood.Name = cases.Title(language.Und, cases.NoLower).String(createFood.Name)

	if createFood.IsAvail < 0 {
		createFood.IsAvail = 0
	} else if createFood.IsAvail > 1 {
		createFood.IsAvail = 1
	} else if createFood.IsAvail != 0 && createFood.IsAvail != 1 {
		createFood.IsAvail = 1
	}

	if createFood.IsVeg < 0 {
		createFood.IsVeg = 0
	} else if createFood.IsVeg > 1 {
		createFood.IsVeg = 1
	} else if createFood.IsVeg != 0 && createFood.IsVeg != 1 {
		createFood.IsVeg = 1
	}

	return nil
}

func validateUpdateFoodReq(updateFood *dto.Food) error {

	if updateFood.ID <= 0 {
		return errors.New(internal.InvalidFoodID)
	} else if updateFood.CategoryID <= 0 {
		return errors.New(internal.InvalidCategoryID)
	} else if updateFood.Price <= 0 {
		return errors.New(internal.InvalidPrice)
	} else if len(updateFood.Name) < 2 {
		return errors.New(internal.InvalidFoodName)
	}

	if updateFood.IsAvail < 0 {
		updateFood.IsAvail = 0
	} else if updateFood.IsAvail > 1 {
		updateFood.IsAvail = 1
	} else if updateFood.IsAvail != 0 && updateFood.IsAvail != 1 {
		updateFood.IsAvail = 1
	}

	if updateFood.IsVeg < 0 {
		updateFood.IsVeg = 0
	} else if updateFood.IsVeg > 1 {
		updateFood.IsVeg = 1
	} else if updateFood.IsVeg != 0 && updateFood.IsVeg != 1 {
		updateFood.IsVeg = 1
	}
	return nil
}

func validateAddOrderItemReq(addOrederItem *dto.OrderItem) error {

	if addOrederItem.ID <= 0 {
		return errors.New(internal.InvalidFoodID)
	} else if addOrederItem.Quantity <= 0 {
		return errors.New(internal.InvalidQuantity)
	}
	return nil
}

func validateCreateInvoice(req *dto.InvoiceCreation) error {
	payMethod := strings.ToLower(req.PaymentMethod)
	if paymentMethodTypes[payMethod] != 1 {
		return errors.New(internal.InvalidPaymentMethod)
	} else if len(req.Location) < 2 {
		return errors.New(internal.InvalidLocation)
	}
	req.Location = cases.Title(language.Und, cases.NoLower).String(req.Location)
	return nil
}

func validateUpdateDelivery(req *dto.DeliveryUpdateRequst) error {
	deliverystat := strings.ToLower(req.Status)
	if req.UserID <= 0 {
		return errors.New(internal.InvalidUserID)
	} else if deliveryStatus[deliverystat] != 1 {
		return errors.New(internal.InvalidDeliveryStatus)
	}
	if req.EndTime == "" {
		return nil
	}
	layout := "2006-01-02 15:04:05"
	_, err := time.Parse(layout, req.EndTime)
	if err != nil {
		log.Println("Invalid time format:", err)
		return errors.New(internal.InvalidTimeFormat)
	} else {
		log.Println("Valid time format")

	}
	return nil

}

func validateCreateUserReq(user *dto.UserSignUpRequest) error {
	if len(user.Firstname) < 2 && containsOnlyAlphabets(user.Firstname) {
		return errors.New(internal.InvalidFirstName)
	} else if len(user.Lastname) < 2 && containsOnlyAlphabets(user.Firstname) {
		return errors.New(internal.InvalidLastName)
	} else if len(user.Phone) < 10 && containsOnlyDigits(user.Phone) {
		return errors.New(internal.InvalidPhoneNumber)
	} else if roleOfUser[strings.ToLower(user.Role)] != 1 {
		return errors.New(internal.InvalidRole)
	} else if !(isValidEmail(user.Email)) {
		return errors.New(internal.InvalidEmail)
	} else if !(isValidPassword(user.Password)) {
		return errors.New(internal.InvalidPassword)
	}

	user.Lastname = cases.Title(language.Und, cases.NoLower).String(user.Lastname)
	user.Firstname = cases.Title(language.Und, cases.NoLower).String(user.Firstname)
	return nil
}
func validateUpdateUserInfo(user dto.UpdateUserInfo) error {
	if len(user.Firstname) < 2 && containsOnlyAlphabets(user.Firstname) {
		return errors.New(internal.InvalidFirstName)
	} else if len(user.Lastname) < 2 && containsOnlyAlphabets(user.Firstname) {
		return errors.New(internal.InvalidLastName)
	} else if len(user.Phone) < 10 && containsOnlyDigits(user.Phone) {
		return errors.New(internal.InvalidPhoneNumber)
	} else if !(isValidEmail(user.Email)) {
		return errors.New(internal.InvalidEmail)
	} else if !(isValidPassword(user.Password)) {
		return errors.New(internal.InvalidPassword)
	}

	user.Lastname = cases.Title(language.Und, cases.NoLower).String(user.Lastname)
	user.Firstname = cases.Title(language.Und, cases.NoLower).String(user.Firstname)
	return nil
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

func containsOnlyAlphabets(s string) bool {

	alphaRegex := regexp.MustCompile("^[a-zA-Z]+$")
	return alphaRegex.MatchString(s)
}
func containsOnlyDigits(s string) bool {
	digitRegex := regexp.MustCompile("^[0-9]+$")
	return digitRegex.MatchString(s)
}
