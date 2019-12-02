package utils

//AppError Is an type of errors that can happen in application
type AppError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
}
