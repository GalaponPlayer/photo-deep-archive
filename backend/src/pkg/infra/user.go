package infra

import (
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
)

type UserRepositoryInfra struct {
}

func NewUserRepositoryInfra() repository.UserRepository {
	return &UserRepositoryInfra{}
}

func (u *UserRepositoryInfra) Create(req *gateway.CreateUserRequest) error {
	//infra
	return nil
}
