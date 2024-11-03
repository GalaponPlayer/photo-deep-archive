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
	LoginID     *string `json:"login_id,omitempty"`
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

	if req.LoginID == nil {
		return errorhandle.NewRequiredButNotFoundError("login_id")
	} else if len(*req.LoginID) == 0 {
		return errorhandle.NewRequiredButNotFoundError("login_id")
	}

	if req.Password == nil {
		return errorhandle.NewRequiredButNotFoundError("password")
		//TODO: cognitoのパスワードルールに合わせる
	} else if len(*req.Password) == 0 {
		return errorhandle.NewRequiredButNotFoundError("password")
	}
	return nil
}

func (req CreateUserUseCaseRequest) ToGateway(id entity.UserID, ts int64) *gateway.CreateUserRequest {
	return gateway.NewCreateUserRequest(
		*entity.NewUser(id, *req.Name, ts),
		*req.MailAddress,
		*req.LoginID,
		*req.Password,
	)
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

	if err := req.Validate(); err != nil {
		return nil, errorhandle.Wrap("req.Validate()", err)
	}

	id, err := lib.GenerateUUIDv4()
	if err != nil {
		return nil, errorhandle.Wrap("lib.GenerateUUIDv4()", err)
	}
	//TODO: IDが存在しているかチェックORフロントで確認できるようなAPI

	createReq := req.ToGateway(entity.UserID(id), lib.GetNowUnixTimeSeconds())
	if err := usecase.userRepository.Create(createReq); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errorhandle.Wrap("userRepository.Create()", err)
	}

	return &CreateUserUseCaseResponse{
		ID: string(id),
	}, nil
}
