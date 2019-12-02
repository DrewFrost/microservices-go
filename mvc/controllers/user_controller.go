package controllers

import (
	"encoding/json"
	"fmt"
	"microservices-go/mvc/services"
	"microservices-go/mvc/utils"
	"net/http"
	"strconv"
)

//GetUserHandler Gets info about user
func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	userIDParam := req.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiErr := &utils.AppError{
			Message: "Invalid User ID",
			Status:  http.StatusBadRequest,
			Code:    "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUserService(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.Status)
		res.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)
	fmt.Println(jsonValue)
	res.Write(jsonValue)

}
