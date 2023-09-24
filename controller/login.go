package controller

import (
	"github.com/biswajitpain/golang-gin-api/dto"
	"github.com/biswajitpain/golang-gin-api/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

// Login implements LoginController.
func (ctlr *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := ctlr.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return ctlr.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}

func NewLoginController(loginService service.LoginService, jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jWtService,
	}
}
