package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service-app/internals/api"
	"user-service-app/internals/repository"
	"user-service-app/internals/service"
)

func main() {

	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)
	handler := api.NewUserHandler(userService)
	handler.RegisterRoutes()
	
	

	fmt.Println("Servern running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}