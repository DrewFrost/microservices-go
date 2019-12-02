package services

import (
	"microservices-go/mvc/domain"
	"microservices-go/mvc/utils"
)

// GetUserService Send request through domain to DB to get user
func GetUserService(userID int64) (*domain.User, *utils.AppError) {
	return domain.GetUser(userID)
}
