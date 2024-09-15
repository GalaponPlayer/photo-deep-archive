package gateway

import "app/src/pkg/domain/entity"

type CreateUserRequest struct {
	User     entity.User `json:"user"`
	LoginID  string      `json:"login_id"`
	Password string      `json:"password"`
}

func NewCreateUserRequest(user entity.User, loginID, password string) *CreateUserRequest {
	return &CreateUserRequest{
		User:     user,
		LoginID:  loginID,
		Password: password,
	}
}
