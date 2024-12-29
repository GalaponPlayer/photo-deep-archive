package repository

import (
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
)

type UserRepository interface {
	// FindAll() ([]*entity.User, error)
	Find(req *gateway.FindUserRequest) (user *entity.User, isNotFound bool, err error)
	Get(id entity.UserID) (user *entity.User, isNotFound bool, err error)
	Create(user *gateway.CreateUserRequest) (*gateway.CreateUserResponse, error)
	// Update(user *entity.User) error
	// Delete(id entity.UserID) error
}
