package infra

import (
	"app/src/pkg/config"
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
	"app/src/pkg/errorhandle"
	"app/src/pkg/infra/auth"
	"app/src/pkg/infra/db"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryInfra struct {
	config        *config.ConfigVariables
	cognitoClient *auth.CognitoClient
	gormClient    *db.GormClient
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
	gormClient, err := db.NewGormClient(*config)
	if err != nil {
		err = errorhandle.Wrap("infra.NewUserRepositoryInfra().db.NewGormClient", err)
		return nil, err
	}

	return &UserRepositoryInfra{
		config:        config,
		cognitoClient: cognitoClient,
		gormClient:    gormClient,
	}, nil
}

func (repo UserRepositoryInfra) Find(req *gateway.FindUserRequest) (user *entity.User, isNotFound bool, err error) {
	user = &entity.User{}
	result := repo.gormClient.DB.Where("mail_address = ?", req.Email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, true, nil
		}
		return nil, false, errorhandle.Wrap("infra.UserRepositoryInfra.Find()", result.Error)
	}

	return user, false, nil
}

func (repo UserRepositoryInfra) Get(id entity.UserID) (user *entity.User, isNotFound bool, err error) {
	user = &entity.User{ID: id}
	result := repo.gormClient.DB.First(&user, id.Value())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, true, nil
		}
		return nil, false, errorhandle.Wrap("infra.UserRepositoryInfra.Get()", result.Error)
	}

	return user, false, nil
}

func (repo UserRepositoryInfra) Create(req *gateway.CreateUserRequest) (*entity.UserID, error) {
	//gorm
	user := req.User
	result := repo.gormClient.DB.Create(&user)
	if result.Error != nil {
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", result.Error)
	}
	//TiDBのAUTO_RANDOM
	userID := user.ID

	//cognito
	isUserConfirmed, err := repo.cognitoClient.SignUp(*req.ToCognitoSignUpInput(repo.config.Cognito.AppClientID, userID.Value()))
	if err != nil {
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", err)
	}
	if !isUserConfirmed {
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("User is not confirmed"))
	}
	return &userID, nil
}
