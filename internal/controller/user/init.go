package user

import (
	"github.com/gin-gonic/gin"
	userService "github.com/kevinnaserwan/coursphere-api/internal/service/user"
	"github.com/kevinnaserwan/coursphere-api/internal/util/jwt"
)

type userController struct {
	UserService userService.Service
}

func NewUserController(router *gin.RouterGroup, userService userService.Service, jwtManager *jwt.JWTManager) {

	controller := &userController{
		UserService: userService,
	}

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetByID)
	router.PUT("/users/:id", controller.Update)
	router.DELETE("/users/:id", controller.Delete)
}
