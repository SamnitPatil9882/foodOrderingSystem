package internal

import "net/http"

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

var (
	InternalServerError   = "Internal server error"
	InvalidCategoryName   = "Invalid category name"
	InvalidCategoryID     = "Invalid id"
	InvalidFoodName       = "Invalid food name"
	InvalidFoodID         = "Invalid food id"
	InvalidPrice          = "Invalid price"
	InvalidQuantity       = "Invalid Quantity"
	InvalidPaymentMethod  = "Invalid Payment Methods"
	InvalidLocation       = "Invalid Location information"
	InvalidUserID         = "Invalid User ID"
	InvalidDeliveryStatus = "Invalid Delivery Stutus"
	InvalidTimeFormat     = "Invalid Time format"
	InvalidFirstName      = "Invalid First Name"
	InvalidLastName       = "Invalid Last Name"
	InvalidPhoneNumber    = "Invalid Phone number"
	InvalidRole           = "Invalid Role"
	InvalidEmail          = "Invalid Email"
	InvalidPassword       = "Invalid Password"
	EmptyCart             = "Cart is Empty"
	Unauthorized          = "login again"
	InvalidID             = "invalid id"
	UnauthorizedAccess    = "Unauthorized access"
	UserAvail             = "User is avail"
	InvalidCredentials    = "eneter valid credentials"
)

const (
	ErrCategoryNotFound                 = CustomError("err:Category not found")
	ErrFailedToFetchCategories          = CustomError("err:Failed to fetch categories")
	ErrFailedToFetchCategory            = CustomError("err:Failed to fetch category")
	ErrFailedToInsertCategory           = CustomError("err:Failed to insert category")
	ErrFailedToUpdateCategory           = CustomError("err:Failed to update category")
	ErrNoRowsAffected                   = CustomError("err:No rows affected")
	ErrEmailExists                      = CustomError("err:email already exist")
	ErrPhoneExists                      = CustomError("err:phone number already exist")
	ErrUserNotFound                     = CustomError("err:user not found")
	ErrEmailOrPasswordIncorrect         = CustomError("err:email or password is incorrect")
	ErrCategoryNameExists               = CustomError("err:category name already exists")
	ErrFoodNameExists                   = CustomError("err:food name already exists")
	ErrFoodNotFound                     = CustomError("err:food not found")
	ErrOrderIdNotExists                 = CustomError("err:order does not exist")
	ErrOrderNotFound                    = CustomError("err:order not found")
	ErrDeliveryBoyIdNotExists           = CustomError("err:user ID does not belong to a delivery boy")
	ErrPrepareDeliveryUpdateStmt        = CustomError("err:failed to prepare delivery update statement")
	ErrExecuteDeliveryUpdateStmt        = CustomError("err:failed to execute delivery update statement")
	ErrNoRowsAffectedAfterUpdate        = CustomError("err:no rows affected after delivery update")
	ErrFailedToFetchDeliveryStatus      = CustomError("err:failed to fetch current delivery status")
	ErrInvalidDeliveryStatusToPickup    = CustomError("err:cannot transition to 'pickup' from current status")
	ErrInvalidDeliveryStatusToDelivered = CustomError("err:cannot transition to 'delivered' from current status")
	ErrInvalidDeliveryId                = CustomError("err:delivery id is not a valid delivery")
	ErrCartIsEmpty = CustomError("err:cart is empty")
	ErrCartOrderItemIdNotFound = CustomError("err:cart order item id is not found")
)


func GetHTTPStatusCode(err error) int {
	switch err {
	case ErrCategoryNotFound,ErrCartOrderItemIdNotFound,ErrCartIsEmpty:
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