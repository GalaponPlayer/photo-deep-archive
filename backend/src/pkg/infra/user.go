package infra

import (
	"app/src/pkg/config"
	"app/src/pkg/domain/entity"
	"app/src/pkg/domain/gateway"
	"app/src/pkg/domain/repository"
	"app/src/pkg/errorhandle"
	"app/src/pkg/infra/auth"
	"app/src/pkg/infra/db"
	"app/src/pkg/lib"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
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
	result := repo.gormClient.DB.Where("email = ?", req.Email).First(&user)
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

func (repo UserRepositoryInfra) Create(req *gateway.CreateUserRequest) (*gateway.CreateUserResponse, error) {
	res := &gateway.CreateUserResponse{}
	//gorm
	user := req.User
	//unique制約のエラーチェックができないため、Findで確認
	findReq := &gateway.FindUserRequest{Email: req.MailAddress}
	findRes, isNotFound, err := repo.Find(findReq)
	if !isNotFound {
		lib.LogInfo("already exists", *findRes)
		res.IsEmailAlreadyExistsError = true
		return res, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("Mail address is already used"))
	}

	result := repo.gormClient.DB.Create(&user)
	if result.Error != nil {
		lib.LogError("UserRepositoryInfra.Create()", result.Error.Error())
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", result.Error)
	}
	//TiDBのAUTO_RANDOMでIDが発行される
	res.ID = user.ID

	//cognito
	isUserConfirmed, err := repo.cognitoClient.SignUp(*req.ToCognitoSignUpInput(repo.config.Cognito.AppClientID, user.ID.Value()))
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		var usernameExists *types.UsernameExistsException
		if errors.As(err, &invalidPassword) {
			repo.gormClient.DB.Delete(&user)
			res.IsPasswordInvalidError = true
			return res, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("Password is invalid"))
		}
		if errors.As(err, &usernameExists) {
			repo.gormClient.DB.Delete(&user)
			res.IsEmailAlreadyExistsError = true
			return res, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("Mail address is already used"))
		}
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", err)
	}
	if !isUserConfirmed {
		return nil, errorhandle.Wrap("infra.UserRepositoryInfra.Create()", errorhandle.NewError("User is not confirmed"))
	}
	return res, nil
}
