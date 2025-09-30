package handler

import (
	"go-users-simple-crud/handler/dto"
	"net/http"
)

func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		response := dto.HealthResponse{
			Status: "UP",
		}
		WriteJSON(w, response, http.StatusOK)
	}
}
