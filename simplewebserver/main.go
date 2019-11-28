package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Web server runnin on localhost:9090")
}
