package internal

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
    ErrCategoryNotFound   = CustomError("Category not found")
    ErrFailedToFetchCategories = CustomError("Failed to fetch categories")
    ErrFailedToFetchCategory = CustomError("Failed to fetch category")
    ErrFailedToInsertCategory = CustomError("Failed to insert category")
    ErrFailedToUpdateCategory = CustomError("Failed to update category")
    ErrNoRowsAffected     = CustomError("No rows affected")
)