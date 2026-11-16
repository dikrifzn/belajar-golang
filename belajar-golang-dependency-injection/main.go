package main

import (
	"belajar-golang-dependency-injection/helper"
	"belajar-golang-dependency-injection/middleware"
	"fmt"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}
func main() {
	server := InitializeServer()

	fmt.Println("Server is running on http://localhost:8080")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
