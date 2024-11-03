package auth

import (
	"app/src/pkg/errorhandle"
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoClient struct {
	Client *cognitoidentityprovider.Client
}

func NewCognitoClient() (*CognitoClient, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, errorhandle.Wrap("infra.NewCognitoClient().config.LoadDefaultConfig", err)
	}
	cognitoClient := cognitoidentityprovider.NewFromConfig(sdkConfig)
	c := &CognitoClient{
		Client: cognitoClient,
	}

	return c, nil
}

func (c *CognitoClient) SignUp(input cognitoidentityprovider.SignUpInput) (bool, error) {
	confirmed := false
	output, err := c.Client.SignUp(context.TODO(), &input)
	if err != nil {
		// //TODO:専用のレスポンス
		// var invalidPassword *types.InvalidPasswordException
		// if errors.As(err, &invalidPassword) {
		// 	errorhandle.Wrap("infra.CognitoClient.SignUp()", err)
		// } else {
		// 	errorhandle.Wrap("infra.CognitoClient.SignUp()", err)
		// }
		return false, errorhandle.Wrap("infra.CognitoClient.SignUp()", err)
	} else {
		confirmed = output.UserConfirmed
	}

	return confirmed, nil
}
