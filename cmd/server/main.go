package main

import (
	"user-service-app/internals/api"
	"user-service-app/internals/repository"
	"user-service-app/internals/service"
)

func main() {

	repo := repository.NewUserRepository()
	userService := service.NewUserService(repo)

	api.RuntMenu(userService)


}