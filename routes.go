package main

import (
	"go-users-simple-crud/middleware"
	"net/http"
)

func initializeRoutes(container *Container) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{id}", container.UserHandler.GetUser())
	mux.HandleFunc("POST /user", container.UserHandler.RegisterUser())
	mux.HandleFunc("DELETE /user/{id}", container.UserHandler.DeleteUser())
	mux.HandleFunc("PUT /user/{id}", container.UserHandler.UpdateUser())
	mux.HandleFunc("GET /user/count", container.UserHandler.CountUsers())

	return middleware.RequestResponseLogger(mux)
}
