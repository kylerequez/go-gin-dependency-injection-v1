package controller

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/middleware"
	"github.com/kylerequez/go-gin-dependency-injection-v1/api/service"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{us: us}
}

func (uc *UserController) InitRoutes(r *gin.Engine, m *middleware.Middleware) {
	apiName := "users"
	prefix := os.Getenv("API_PREFIX")

	v1 := r.Group(prefix + apiName)
	{
		v1.GET("/", uc.GetAllUsers)
		v1.GET("/:id", uc.GetUserById)
		v1.Use(m.UserAuthJWTMiddleware())
		v1.POST("/", uc.CreateUser)
		v1.PATCH("/:id", uc.PatchUpdateUser)
		v1.PUT("/:id", uc.PutUpdateUser)
		v1.DELETE("/:id", uc.DeleteUser)
	}

	log.Println(":::-::: Initialized User Routes")
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	uc.us.GetAllUsers(c)
}

func (uc *UserController) GetUserById(c *gin.Context) {
	uc.us.GetUserById(c)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	uc.us.CreateUser(c)
}

func (uc *UserController) PatchUpdateUser(c *gin.Context) {
	uc.us.PatchUpdateUser(c)
}

func (uc *UserController) PutUpdateUser(c *gin.Context) {
	uc.us.PutUpdateUser(c)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	uc.us.DeleteUser(c)
}
