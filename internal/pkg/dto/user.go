package dto

type User struct {
	ID        int    `json:"id"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}

type UserSignUpRequest struct {
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID        int    `json:"id"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
}
type UserLoginResponse struct {
	ID   int
	Role string
}
type UpdateUserInfo struct {
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}
