package gateway

import "app/src/pkg/domain/entity"

type CreateUserRequest struct {
	User        entity.User `json:"user"`
	MailAddress string      `json:"mail_address"`
	LoginID     string      `json:"login_id"`
	Password    string      `json:"password"`
}

func NewCreateUserRequest(user entity.User, mailAddress string, loginID string, password string) *CreateUserRequest {
	return &CreateUserRequest{
		User:        user,
		MailAddress: mailAddress,
		LoginID:     loginID,
		Password:    password,
	}
}
