package gateway

//TODO: gatewayではなくportにする

import (
	"app/src/pkg/domain/entity"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CreateUserRequest struct {
	User        entity.User `json:"user"`
	MailAddress string      `json:"mail_address"`
	Password    string      `json:"password"`
}

func NewCreateUserRequest(user entity.User, mailAddress string, password string) *CreateUserRequest {
	return &CreateUserRequest{
		User:        user,
		MailAddress: mailAddress,
		Password:    password,
	}
}

func (req CreateUserRequest) ToCognitoSignUpInput(cognitoAppClientId string, userId uint) *cognitoidentityprovider.SignUpInput {
	userIdStr := strconv.FormatUint(uint64(userId), 10) // uint を string に変換

	return &cognitoidentityprovider.SignUpInput{
		ClientId: &cognitoAppClientId,
		Username: &req.MailAddress,
		Password: &req.Password,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(req.MailAddress)},
			{Name: aws.String("custom:id"), Value: aws.String(userIdStr)}, // string に変換した userId を使用
		},
	}
}

type FindUserRequest struct {
	Email string `json:"email"`
}
