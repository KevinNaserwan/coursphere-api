package course

import (
	"github.com/gin-gonic/gin"
	courseService "github.com/kevinnaserwan/coursphere-api/internal/service/course"
)

type courseController struct {
	CourseService courseService.Service
}

func NewCourseController(router *gin.RouterGroup, courseService courseService.Service) {
	controller := courseController{
		CourseService: courseService,
	}

	router.POST("/create", controller.Create)
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
	router.POST("/add-video/:id", controller.AddCourseVideos)
}
