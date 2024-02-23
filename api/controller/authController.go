package controller

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/middleware"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/service"
)

type AuthController struct {
	as *service.AuthService
}

func NewAuthController(as *service.AuthService) *AuthController {
	return &AuthController{as: as}
}

func (ac *AuthController) InitRoutes(r *gin.Engine, m *middleware.Middleware) {
	apiName := "auth"
	prefix := os.Getenv("API_PREFIX")

	v1 := r.Group(prefix + apiName)
	{
		v1.POST("/register", ac.RegistrationHandler)
		v1.POST("/login", ac.LoginHandler)
	}
	log.Println(":::-::: Initialized Auth Routes")
}

func (ac *AuthController) RegistrationHandler(c *gin.Context) {
	ac.as.RegistrationHandler(c)
}

func (ac *AuthController) LoginHandler(c *gin.Context) {
	ac.as.LoginHandler(c)
}
