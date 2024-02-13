package dto
type ErrorResponse struct{
	Error int `json:"error"`
	Description  string `json:"description"`
}