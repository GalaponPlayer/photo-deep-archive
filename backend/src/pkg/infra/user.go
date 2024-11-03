package infra

import (
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
	"app/src/pkg/errorhandle"
	"app/src/pkg/infra/auth"
)

type UserRepositoryInfra struct {
}

func NewUserRepositoryInfra() repository.UserRepository {
	return &UserRepositoryInfra{}
}

func (u *UserRepositoryInfra) Create(req *gateway.CreateUserRequest) error {
	//TODO: cognito
	cognitoClient, err := auth.NewCognitoClient()
	isUserConfirmed, err := cognitoClient.SignUp(*req.ToCognitoSignUpInput())
	if err != nil {
		return errorhandle.Wrap("infra.UserRepositoryInfra.Create()", err)
	}
	if !isUserConfirmed {
		return errorhandle.New("infra.UserRepositoryInfra.Create()", "User is not confirmed")
	}
	//TODO: cockroachdb
	return nil
}
