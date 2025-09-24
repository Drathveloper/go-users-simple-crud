package main

import (
	"go-users-simple-crud/handler"
	"go-users-simple-crud/repository"
	"go-users-simple-crud/service"
)

type Container struct {
	UserHandler            *handler.User
	UserService            *service.User
	UserInMemoryRepository *repository.UserInMemory
}

func initializeDependencies() *Container {
	container := Container{}
	container.UserInMemoryRepository = repository.NewUserInMemoryRepository()
	container.UserService = service.NewUserService(container.UserInMemoryRepository)
	container.UserHandler = handler.NewUserHandler(container.UserService)
	return &container
}
