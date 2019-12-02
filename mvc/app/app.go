package app

import (
	"microservices-go/mvc/controllers"
	"net/http"
)

//StartApp Starts Web App
func StartApp() {
	http.HandleFunc("/users", controllers.GetUserHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
