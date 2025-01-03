package main

import (
	"app/src/pkg/factory"
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func setupRouter() *gin.Engine {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := gin.Default()
	setCors(router)

	v1Router := router.Group("/v1")
	v1Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "CICD Success!",
		})
	})

	// user
	v1Router.POST("/user", factory.CreateUser)

	return router
}

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := setupRouter()
	ginLambda = ginadapter.New(router)
}

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8010"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-Amz-Date", "Authorization", "X-Api-Key", "X-Amz-Security-Token", "X-Amz-User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	log.Printf("req.Path: " + req.Path)
	log.Printf("req: %v", req)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// 環境変数で実行モードを判定
	if os.Getenv("GIN_MODE") == "debug" {
		// ローカル開発モード
		router := setupRouter()
		log.Println("Running in local debug mode...")
		router.Run(":8080") // ローカルサーバーをポート8080で起動
	} else {
		// Lambdaモード
		log.Println("Running in Lambda mode...")
		lambda.Start(Handler)
	}
}
