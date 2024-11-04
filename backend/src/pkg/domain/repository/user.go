package repository

import (
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
)

type UserRepository interface {
	// FindAll() ([]*entity.User, error)
	// FindByID(id entity.UserID) (*entity.User, error)
	Create(user *gateway.CreateUserRequest) (*entity.UserID, error)
	// Update(user *entity.User) error
	// Delete(id entity.UserID) error
}
