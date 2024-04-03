package internal

import "net/http"

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

var (
	InternalServerError     = "Internal server error"
	InvalidCategoryName     = "Invalid category name"
	InvalidCategoryID       = "Invalid category id"
	InvalidFoodName         = "Invalid food name"
	InvalidFoodID           = "Invalid food id"
	InvalidPrice            = "Invalid price"
	InvalidQuantity         = "Invalid Quantity"
	InvalidPaymentMethod    = "Invalid Payment Methods"
	InvalidLocation         = "Invalid Location information"
	InvalidUserID           = "Invalid User ID"
	InvalidDeliveryStatus   = "Invalid Delivery Stutus"
	InvalidTimeFormat       = "Invalid Time format"
	InvalidFirstName        = "Invalid First Name"
	InvalidLastName         = "Invalid Last Name"
	InvalidPhoneNumber      = "Invalid Phone number"
	InvalidRole             = "Invalid Role"
	InvalidEmail            = "Invalid Email"
	InvalidPassword         = "Invalid Password"
	EmptyCart               = "Cart is Empty"
	Unauthorized            = "login again"
	InvalidID               = "invalid id"
	UnauthorizedAccess      = "Unauthorized access"
	UserAvail               = "User is avail"
	InvalidCredentials      = "eneter valid credentials"
	InvalidCategoryIsActive = "Invalid category is active"
	InvalidIsAvail          = "Invalid isavail"
	InvalidIsVeg            = "Invalid isveg"
	InvalidDescription      = "Invalid description"
	InvalidImgUrl           = "Invalid imgurl"
)

const (
	ErrCategoryNotFound                 = CustomError("Category not found")
	ErrFailedToFetchCategories          = CustomError("Failed to fetch categories")
	ErrFailedToFetchCategory            = CustomError("Failed to fetch category")
	ErrFailedToInsertCategory           = CustomError("Failed to insert category")
	ErrFailedToUpdateCategory           = CustomError("Failed to update category")
	ErrNoRowsAffected                   = CustomError("No rows affected")
	ErrEmailExists                      = CustomError("email already exist")
	ErrPhoneExists                      = CustomError("phone number already exist")
	ErrUserNotFound                     = CustomError("user not found")
	ErrEmailOrPasswordIncorrect         = CustomError("email or password is incorrect")
	ErrCategoryNameExists               = CustomError("category name already exists")
	ErrFoodNameExists                   = CustomError("food name already exists")
	ErrFoodNotFound                     = CustomError("food not found")
	ErrOrderIdNotExists                 = CustomError("order does not exist")
	ErrOrderNotFound                    = CustomError("order not found")
	ErrDeliveryBoyIdNotExists           = CustomError("user ID does not belong to a delivery boy")
	ErrPrepareDeliveryUpdateStmt        = CustomError("failed to prepare delivery update statement")
	ErrExecuteDeliveryUpdateStmt        = CustomError("failed to execute delivery update statement")
	ErrNoRowsAffectedAfterUpdate        = CustomError("no rows affected after delivery update")
	ErrFailedToFetchDeliveryStatus      = CustomError("failed to fetch current delivery status")
	ErrInvalidDeliveryStatusToPickup    = CustomError("cannot transition to 'pickup' from current status")
	ErrInvalidDeliveryStatusToDelivered = CustomError("cannot transition to 'delivered' from current status")
	ErrInvalidDeliveryId                = CustomError("delivery id is not a valid delivery")
	ErrCartIsEmpty                      = CustomError("cart is empty")
	ErrCartOrderItemIdNotFound          = CustomError("cart order item id is not found")
	ErrFailedToDeleteCategory           = CustomError("failed to delete category")
	ErrInvalidCategoryId                = CustomError("invalid category id enter valid category id")

)

func GetHTTPStatusCode(err error) int {
	switch err {
	case ErrCategoryNotFound, ErrCartOrderItemIdNotFound, ErrCartIsEmpty:
		return http.StatusNotFound
	case ErrFailedToFetchCategories, ErrFailedToFetchCategory, ErrFailedToInsertCategory,
		ErrFailedToUpdateCategory, ErrNoRowsAffected, ErrPrepareDeliveryUpdateStmt,
		ErrExecuteDeliveryUpdateStmt, ErrNoRowsAffectedAfterUpdate, ErrFailedToFetchDeliveryStatus:
		return http.StatusInternalServerError
	case ErrEmailExists, ErrPhoneExists, ErrCategoryNameExists, ErrFoodNameExists:
		return http.StatusConflict
	case ErrUserNotFound, ErrFoodNotFound, ErrOrderIdNotExists, ErrOrderNotFound,
		ErrDeliveryBoyIdNotExists, ErrInvalidDeliveryId:
		return http.StatusNotFound
	case ErrEmailOrPasswordIncorrect:
		return http.StatusUnauthorized
	case ErrInvalidDeliveryStatusToPickup, ErrInvalidDeliveryStatusToDelivered:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
