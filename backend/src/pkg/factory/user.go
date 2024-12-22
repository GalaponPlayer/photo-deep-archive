package factory

import (
	"app/src/pkg/handler"
	"app/src/pkg/infra"
	"app/src/pkg/lib"
	"app/src/pkg/usecase"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	//usecase
	userRepositoryInfra, err := infra.NewUserRepositoryInfra()
	if err != nil {
		if err != nil {
			lib.LogError("Process Error:", err)
			// c.JSON(500, gin.H{"error": err.Error()})
			c.JSON(500, lib.ErrorResponseBody{Message: "init repository error"})
			return
		}
	}
	usecase := usecase.NewCreateUserUseCase(userRepositoryInfra)

	//usecase
	handler := handler.NewCreateUserHandler(usecase)
	handler.CreateUser(c)
}
