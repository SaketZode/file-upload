package main

import (
	"file-upload/services"
	"fmt"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/upload", services.UploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
