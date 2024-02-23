package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/controller"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/middleware"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/repository"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/service"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	js := service.NewJwtService()
	m := middleware.NewMiddleware(js)

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us)
	uc.InitRoutes(r, m)

	as := service.NewAuthService(ur, js)
	ac := controller.NewAuthController(as)
	ac.InitRoutes(r, m)
}
