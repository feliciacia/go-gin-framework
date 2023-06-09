package controller

import (
	"errors"

	"github.com/feliciacia/go-gin-framework/go-gin-framework/dto"
	"github.com/feliciacia/go-gin-framework/go-gin-framework/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) (string, error)
}

type loginController struct {
	loginService service.LoginServiceInterface
	jWtService   service.JWTService //
}

func NewLoginController(loginService service.LoginServiceInterface,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) (string, error) {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return "", err
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true), nil
	}
	return "", errors.New("is not authenticated")
}
