package mapper

import (
	"go-users-simple-crud/model"
	"go-users-simple-crud/repository/entity"
)

func UserEntityToUserModel(entity entity.User) model.User {
	return model.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		BirthDate: entity.BirthDate,
	}
}

func UserModelToUserEntity(model model.User) entity.User {
	return entity.User{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		BirthDate: model.BirthDate,
	}
}
