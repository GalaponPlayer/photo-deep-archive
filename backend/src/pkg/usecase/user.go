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

// TODO: Httpリクエストの全てを受け取るべきでは？
// 必要なものをuseaCaseに渡す？　それを整形するのはhandlerの役割？
type CreateUserUseCaseRequest struct {
	Name        *string `json:"name,omitempty"`
	MailAddress *string `json:"mail_address,omitempty"`
	Password    *string `json:"password,omitempty"`
}

func (req CreateUserUseCaseRequest) Validate() error {
	if req.Name == nil {
		return errorhandle.NewRequiredButNotFoundError("name")
	} else if len(*req.Name) == 0 {
		return errorhandle.NewRequiredButNotFoundError("name")
	}

	if req.MailAddress == nil {
		return errorhandle.NewRequiredButNotFoundError("mail_address")
	} else if len(*req.MailAddress) == 0 {
		return errorhandle.NewRequiredButNotFoundError("mail_address")
	}

	if req.Password == nil {
		return errorhandle.NewRequiredButNotFoundError("password")
		//TODO: cognitoのパスワードルールに合わせる
	} else if len(*req.Password) == 0 {
		return errorhandle.NewRequiredButNotFoundError("password")
	}
	return nil
}

func (req CreateUserUseCaseRequest) ToGateway() *gateway.CreateUserRequest {
	return gateway.NewCreateUserRequest(
		*entity.NewUser(entity.UserID(0), *req.Name, *req.MailAddress),
		*req.MailAddress,
		*req.Password,
	)
}

type CreateUserUseCaseResponse struct {
	ID                        entity.UserID `json:"id"`
	IsEmailAlreadyExistsError bool          `json:"is_email_already_exists_error"`
	IsPasswordInvalidError    bool          `json:"is_password_invalid_error"`
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

	if err := req.Validate(); err != nil {
		return nil, errorhandle.Wrap("req.Validate()", err)
	}

	createReq := req.ToGateway()
	createRes, err := usecase.userRepository.Create(createReq)
	if err != nil {
		lib.LogError("createRes", err)
		if createRes == nil {
			return nil, errorhandle.Wrap("userRepository.Create()", err)
		}
		if createRes.IsEmailAlreadyExistsError {
			return &CreateUserUseCaseResponse{
				IsEmailAlreadyExistsError: true,
			}, nil
		}
		if createRes.IsPasswordInvalidError {
			return &CreateUserUseCaseResponse{
				IsPasswordInvalidError: true,
			}, nil
		}
	}

	return &CreateUserUseCaseResponse{
		ID: createRes.ID,
	}, nil
}
