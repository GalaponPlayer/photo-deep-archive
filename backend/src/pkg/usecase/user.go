package usecase

import (
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
)

type CreateUserUseCase interface {
	Do(req *CreateUserUseCaseRequest) (*CreateUserUseCaseResponse, error)
}

type CreateUserUseCaseRequest struct {
	Name     string `json:"name"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}

type CreateUserUseCaseResponse struct {
	ID string `json:"id"`
}

type createUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(userRepository repository.UserRepository) CreateUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}

func (usecase createUserUseCase) Do(req *CreateUserUseCaseRequest) (*CreateUserUseCaseResponse, error) {
	user := &entity.User{
		Name: req.Name,
	}
	createReq := gateway.NewCreateUserRequest(*user, req.LoginID, req.Password)
	if err := usecase.userRepository.Create(createReq); err != nil {
		return nil, err
	}
	return &CreateUserUseCaseResponse{
		ID: string(user.ID),
	}, nil
}
