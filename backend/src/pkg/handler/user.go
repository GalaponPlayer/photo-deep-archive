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
		c.JSON(int(lib.StatusCodeBadRequest), lib.ErrorResponseBody{Message: "invalid request"})
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(int(lib.StatusCodeBadRequest), lib.ErrorResponseBody{Message: "invalid request"})
		return
	}

	lib.LogInfo("Request:", req)

	//TODO: logging
	res, err := handler.useCase.Do(req)
	if err != nil {
		c.JSON(int(lib.StatusCodeInternalServerError), lib.ErrorResponseBody{Message: "internal server error"})
		return
	}

	if res.IsEmailAlreadyExistsError {
		c.JSON(int(lib.CreateUserStatusCodeEmailAlreadyExists), lib.ErrorResponseBody{Message: "email already exists"})
		return
	} else if res.IsPasswordInvalidError {
		c.JSON(int(lib.CreateUserStatusCodePasswordInvalid), lib.ErrorResponseBody{Message: "invalid password"})
		return
	}

	lib.LogInfo("Response:", res)

	c.JSON(200, res)
}
