package config

import (
	"app/src/pkg/errorhandle"
	"os"
)

const (
	EnvCognitoUserPoolID  = "COGNITO_USER_POOL_ID"
	EnvCognitoAppClientID = "COGNITO_APP_CLIENT_ID"

	EnvTiDBUser     = "TIDB_USER"
	EnvTiDNHost     = "TIDB_HOST"
	EnvTiDBPort     = "TIDB_PORT"
	EnvTiDBPassword = "TIDB_PASSWORD"
)

type ConfigVariables struct {
	Cognito CognitoConfig
	TiDB    TiDBConfig
}

func NewConfigVariables() (*ConfigVariables, error) {
	cfg := &ConfigVariables{}

	cfg.Cognito.UserPoolID = os.Getenv(EnvCognitoUserPoolID)
	cfg.Cognito.AppClientID = os.Getenv(EnvCognitoAppClientID)
	cfg.TiDB.User = os.Getenv(EnvTiDBUser)
	cfg.TiDB.Host = os.Getenv(EnvTiDNHost)
	cfg.TiDB.Port = os.Getenv(EnvTiDBPort)
	cfg.TiDB.Password = os.Getenv(EnvTiDBPassword)

	if cfg.Cognito.UserPoolID == "" {
		return nil, errorhandle.NewInitializeError("UserPoolID")
	}
	if cfg.Cognito.AppClientID == "" {
		return nil, errorhandle.NewInitializeError("AppClientID")
	}
	if cfg.TiDB.User == "" {
		return nil, errorhandle.NewInitializeError("UserName")
	}
	if cfg.TiDB.Host == "" {
		return nil, errorhandle.NewInitializeError("Host")
	}
	if cfg.TiDB.Port == "" {
		return nil, errorhandle.NewInitializeError("Port")
	}
	if cfg.TiDB.Password == "" {
		return nil, errorhandle.NewInitializeError("Password")
	}

	return cfg, nil
}

type CognitoConfig struct {
	UserPoolID  string
	AppClientID string
}

type TiDBConfig struct {
	User     string
	Host     string
	Port     string
	Password string
}
