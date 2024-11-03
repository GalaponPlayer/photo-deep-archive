package infra

import (
	"app/src/pkg/config"
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
	"app/src/pkg/errorhandle"
	"app/src/pkg/infra/auth"
)

type UserRepositoryInfra struct {
	config        *config.ConfigVariables
	cognitoClient *auth.CognitoClient
}

func NewUserRepositoryInfra() (repository.UserRepository, error) {
	config, err := config.NewConfigVariables()
	if err != nil {
		err = errorhandle.Wrap("infra.NewUserRepositoryInfra().config.NewConfigVariables", err)
		return nil, err
	}
	cognitoClient, err := auth.NewCognitoClient()
	if err != nil {
		err = errorhandle.Wrap("infra.NewUserRepositoryInfra().auth.NewCognitoClient", err)
		return nil, err
	}

	return &UserRepositoryInfra{
		config:        config,
		cognitoClient: cognitoClient,
	}, nil
}

func (repo UserRepositoryInfra) Create(req *gateway.CreateUserRequest) error {
	//TODO: cognito
	isUserConfirmed, err := repo.cognitoClient.SignUp(*req.ToCognitoSignUpInput(repo.config.Cognito.AppClientID))
	if err != nil {
		return errorhandle.Wrap("infra.UserRepositoryInfra.Create()", err)
	}
	if !isUserConfirmed {
		return errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("User is not confirmed"))
	}
	//TODO: cockroachdb
	return nil
}
