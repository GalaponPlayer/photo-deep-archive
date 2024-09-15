package factory

import (
	"app/src/pkg/handler"
	"app/src/pkg/infra"
	"app/src/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//usecase
	userRepositoryInfra := infra.NewUserRepositoryInfra()
	usecase := usecase.NewCreateUserUseCase(userRepositoryInfra)

	//usecase
	handler := handler.NewCreateUserHandler(usecase)
	handler.CreateUser(c)
}
