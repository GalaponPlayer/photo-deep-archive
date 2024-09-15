package main

import (
	"app/src/pkg/factory"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := gin.Default()
	v1Router := router.Group("/v1")
	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CICD Success!",
		})
	})

	// user
	v1Router.POST("/user", factory.CreateUser)

	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	log.Printf("req.Path: " + req.Path)
	log.Printf("req: %v", req)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
