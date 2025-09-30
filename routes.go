package main

import (
	"go-users-simple-crud/handler"
	"go-users-simple-crud/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func initializeRoutes(container *Container) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health())

	mux.HandleFunc("GET /user/{id}", container.UserHandler.GetUser())
	mux.HandleFunc("POST /user", container.UserHandler.RegisterUser())
	mux.HandleFunc("DELETE /user/{id}", container.UserHandler.DeleteUser())
	mux.HandleFunc("PUT /user/{id}", container.UserHandler.UpdateUser())
	mux.HandleFunc("GET /user/count", container.UserHandler.CountUsers())

	mux.Handle("/metrics", promhttp.Handler())

	return middleware.RequestResponseLogger(mux)
}
