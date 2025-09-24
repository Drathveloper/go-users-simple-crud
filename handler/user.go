package handler

import (
	"errors"
	"go-users-simple-crud/handler/dto"
	"go-users-simple-crud/handler/mapper"
	"go-users-simple-crud/model"
	"go-users-simple-crud/service"
	"log/slog"
	"net/http"
	"strconv"
)

type UserService interface {
	GetUser(id int64) (model.User, error)
	RegisterUser(user model.User) (int64, error)
	DeleteUser(id int64) error
	UpdateUser(user model.User) error
	Count() (int, error)
}

const userIDPathParam = "id"

type User struct {
	userService UserService
}

func NewUserHandler(userService UserService) *User {
	return &User{
		userService: userService,
	}
}

func (h *User) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPathParam), 10, 64)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		user, err := h.userService.GetUser(userID)
		if err != nil {
			h.handleUserServiceError(w, err)
			return
		}
		response := mapper.UserToGetUsersResponse(user)
		WriteJSON(w, response, http.StatusOK)
	}
}

func (h *User) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody dto.RegisterUserRequest
		if err := ReadJSON(r, &requestBody); err != nil {
			WriteError(w, err, http.StatusBadRequest)
		}
		user, err := mapper.RegisterUserRequestToUser(requestBody)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		id, err := h.userService.RegisterUser(user)
		if err != nil {
			h.handleUserServiceError(w, err)
			return
		}
		response := mapper.IDToRegisterUserResponse(id)
		WriteJSON(w, response, http.StatusCreated)
	}
}

func (h *User) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPathParam), 10, 64)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		if err = h.userService.DeleteUser(userID); err != nil {
			slog.Warn("delete user failed", "id", userID, "error", err)
		}
		WriteJSON(w, nil, http.StatusNoContent)
	}
}

func (h *User) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.ParseInt(r.PathValue(userIDPathParam), 10, 64)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		var requestBody dto.UpdateUserRequest
		if err = ReadJSON(r, &requestBody); err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		user, err := mapper.UpdateUserRequestAndIDToUser(userID, requestBody)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		if err = h.userService.UpdateUser(user); err != nil {
			h.handleUserServiceError(w, err)
			return
		}
		WriteJSON(w, nil, http.StatusNoContent)
	}
}

func (h *User) CountUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		numUsers, err := h.userService.Count()
		if err != nil {
			WriteError(w, err, http.StatusInternalServerError)
			return
		}
		response := mapper.NumUsersToCountResponse(numUsers)
		WriteJSON(w, response, http.StatusOK)
	}
}

func (h *User) handleUserServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrUserNotFound):
		WriteError(w, err, http.StatusNotFound)
	default:
		WriteError(w, err, http.StatusInternalServerError)
	}
}
