package config

import (
	"app/src/pkg/errorhandle"
	"os"
)

const (
	EnvCognitoUserPoolID  = "COGNITO_USER_POOL_ID"
	EnvCognitoAppClientID = "COGNITO_APP_CLIENT_ID"
)

type ConfigVariables struct {
	Cognito CognitoConfig
}

func NewConfigVariables() (*ConfigVariables, error) {
	cfg := &ConfigVariables{}

	cfg.Cognito.UserPoolID = os.Getenv(EnvCognitoUserPoolID)
	cfg.Cognito.AppClientID = os.Getenv(EnvCognitoAppClientID)

	if cfg.Cognito.UserPoolID == "" {
		return nil, errorhandle.NewInitializeError("UserPoolID")
	}
	if cfg.Cognito.AppClientID == "" {
		return nil, errorhandle.NewInitializeError("AppClientID")
	}

	return cfg, nil
}

type CognitoConfig struct {
	UserPoolID  string
	AppClientID string
}
