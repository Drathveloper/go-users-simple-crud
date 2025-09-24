package repository

import (
	"errors"
	"fmt"
	"go-users-simple-crud/model"
	"go-users-simple-crud/repository/entity"
	"sync"
)

var ErrItemNotFound = errors.New("item not found")

type UserInMemory struct {
	users map[int64]entity.User
	sync.Mutex
}

func NewUserInMemoryRepository() *UserInMemory {
	return &UserInMemory{
		users: make(map[int64]entity.User),
	}
}

func (u *UserInMemory) FindByID(id int64) (model.User, error) {
	u.Lock()
	defer u.Unlock()
	user, ok := u.users[id]
	if !ok {
		return model.User{}, fmt.Errorf("find user by id failed: %w", ErrItemNotFound)
	}
	return model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}, nil
}

func (u *UserInMemory) Save(user model.User) (model.User, error) {
	u.Lock()
	defer u.Unlock()
	id := int64(len(u.users) + 1)
	u.users[id] = entity.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}
	user.ID = id
	return user, nil
}

func (u *UserInMemory) Delete(id int64) error {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.users[id]; !ok {
		return fmt.Errorf("delete user failed: %w", ErrItemNotFound)
	}
	delete(u.users, id)
	return nil
}

func (u *UserInMemory) Update(user model.User) error {
	u.Lock()
	defer u.Unlock()
	if _, ok := u.users[user.ID]; !ok {
		return fmt.Errorf("update user failed: %w", ErrItemNotFound)
	}
	u.users[user.ID] = entity.User{
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
	}
	return nil
}

func (u *UserInMemory) Count() (int, error) {
	u.Lock()
	defer u.Unlock()
	return len(u.users), nil
}
