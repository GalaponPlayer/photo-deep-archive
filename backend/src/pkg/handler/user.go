package handler

import (
	"app/src/pkg/lib"
	"app/src/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type CreateUserHandler interface {
	CreateUser(c *gin.Context)
}

type createUserHandler struct {
	useCase usecase.CreateUserUseCase
}

func NewCreateUserHandler(useCase usecase.CreateUserUseCase) CreateUserHandler {
	return &createUserHandler{
		useCase: useCase,
	}
}

func (handler createUserHandler) CreateUser(c *gin.Context) {
	req := &usecase.CreateUserUseCaseRequest{}
	//TODO: validation
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	lib.LogInfo("Request:", req)

	//TODO: logging
	res, err := handler.useCase.Do(req)
	//TODO: error handling
	//TODO: change status code by res
	if err != nil {
		lib.LogError("Process Error:", err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	lib.LogInfo("Response:", res)

	c.JSON(200, res)
}
