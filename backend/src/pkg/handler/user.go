package handler

import (
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
	if err := c.BindJSON(req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := handler.useCase.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
