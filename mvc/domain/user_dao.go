package domain

import (
	"microservices-go/mvc/utils"
	"net/http"
)

var users = map[int64]*User{
	123: &User{ID: 123, FirstName: "Dron", LastName: "Coldone", Email: "test@com.com"},
}

// GetUser Gets user in db
func GetUser(userID int64) (*User, *utils.AppError) {
	user := users[userID]
	if user == nil {
		return nil, &utils.AppError{
			Message: "User not found",
			Status:  http.StatusNotFound,
			Code:    "not found",
		}
	}
	return user, nil
}
