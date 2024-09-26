package usecase

import (
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
	"app/src/pkg/errorhandle"
	"app/src/pkg/lib"
)

type CreateUserUseCase interface {
	Do(req *CreateUserUseCaseRequest) (*CreateUserUseCaseResponse, error)
}

type CreateUserUseCaseRequest struct {
	Name        string `json:"name"`
	MailAddress string `json:"mail_address"`
	LoginID     string `json:"login_id"`
	Password    string `json:"password"`
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

	id, err := lib.GenerateUUIDv4()
	if err != nil {
		return nil, errorhandle.Wrap("lib.GenerateUUIDv4()", err)
	}

	user := entity.NewUser(entity.UserID(id), req.Name, lib.GetNowUnixTimeSeconds())
	createReq := gateway.NewCreateUserRequest(*user, req.MailAddress, req.LoginID, req.Password)
	//todo: time uuid
	if err := usecase.userRepository.Create(createReq); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errorhandle.Wrap("userRepository.Create()", err)
	}

	return &CreateUserUseCaseResponse{
		ID: string(user.ID),
	}, nil
}
