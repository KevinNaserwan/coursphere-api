package coursecategory

import (
	"github.com/gin-gonic/gin"
	courseCategoryService "github.com/kevinnaserwan/coursphere-api/internal/service/course_category"
)

type courseCategoryController struct {
	CourseCategoryService courseCategoryService.Service
}

func NewCourseCategoryController(router *gin.RouterGroup, courseCategoryService courseCategoryService.Service) {
	controller := courseCategoryController{
		CourseCategoryService: courseCategoryService,
	}

	router.POST("/create", controller.Create)
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}
