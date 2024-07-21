package mentor

import (
	"github.com/gin-gonic/gin"
	mentorService "github.com/kevinnaserwan/coursphere-api/internal/service/mentor"
)

type mentorController struct {
	MentorService mentorService.Service
}

func NewMentorController(router *gin.RouterGroup, mentorService mentorService.Service) {
	controller := mentorController{
		MentorService: mentorService,
	}

	router.POST("/create", controller.Create)
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}
