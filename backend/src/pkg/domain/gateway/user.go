package gateway

import (
	"app/src/pkg/domain/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

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

func (req CreateUserRequest) ToCognitoSignUpInput(cognitoAppClientId string) *cognitoidentityprovider.SignUpInput {
	return &cognitoidentityprovider.SignUpInput{
		ClientId: &cognitoAppClientId,
		Username: &req.LoginID,
		Password: &req.Password,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(req.MailAddress)},
		},
	}
}
