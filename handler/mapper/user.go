package mapper

import (
	"fmt"
	"go-users-simple-crud/handler/dto"
	"go-users-simple-crud/model"
	"time"
)

func UserToGetUsersResponse(user model.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate.Format("02-01-2006"),
	}
}

func RegisterUserRequestToUser(requestBody dto.RegisterUserRequest) (model.User, error) {
	birthDate, err := time.Parse("02-01-2006", requestBody.BirthDate)
	if err != nil {
		return model.User{}, fmt.Errorf("invalid birth date: %w", err)
	}
	return model.User{
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		BirthDate: birthDate,
	}, nil
}

func UpdateUserRequestAndIDToUser(id int64, requestBody dto.UpdateUserRequest) (model.User, error) {
	birthDate, err := time.Parse("02-01-2006", requestBody.BirthDate)
	if err != nil {
		return model.User{}, fmt.Errorf("invalid birth date: %w", err)
	}
	return model.User{
		ID:        id,
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		BirthDate: birthDate,
	}, nil
}

func IDToRegisterUserResponse(id int64) dto.RegisterUserResponse {
	return dto.RegisterUserResponse{
		ID: id,
	}
}

func NumUsersToCountResponse(numUsers int) dto.CountResponse {
	return dto.CountResponse{
		Message: fmt.Sprintf("There are %d users registered", numUsers),
	}
}
