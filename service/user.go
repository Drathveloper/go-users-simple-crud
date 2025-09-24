package service

import (
	"errors"
	"fmt"
	"go-users-simple-crud/model"
	"go-users-simple-crud/repository"
)

var ErrUserNotFound = errors.New("user not found in repository")

type UserRepository interface {
	FindByID(id int64) (model.User, error)
	Save(user model.User) (model.User, error)
	Delete(id int64) error
	Update(user model.User) error
	Count() (int, error)
}

type User struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (s *User) GetUser(id int64) (model.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return model.User{}, s.handleError("find user by id failed", err)
	}
	return user, nil
}

func (s *User) RegisterUser(user model.User) (int64, error) {
	userWithID, err := s.userRepository.Save(user)
	if err != nil {
		return 0, s.handleError("register user failed", err)
	}
	return userWithID.ID, nil
}

func (s *User) DeleteUser(id int64) error {
	if err := s.userRepository.Delete(id); err != nil {
		return s.handleError("delete user failed", err)
	}
	return nil
}

func (s *User) UpdateUser(user model.User) error {
	if err := s.userRepository.Update(user); err != nil {
		return s.handleError("update user failed", err)
	}
	return nil
}

func (s *User) Count() (int, error) {
	numUsers, err := s.userRepository.Count()
	if err != nil {
		return 0, s.handleError("count users failed", err)
	}
	return numUsers, nil
}

func (s *User) handleError(baseErrMsg string, err error) error {
	switch {
	case errors.Is(err, repository.ErrItemNotFound):
		return fmt.Errorf("%s: %w", baseErrMsg, ErrUserNotFound)
	default:
		return fmt.Errorf("%s: %w", baseErrMsg, err)
	}
}
