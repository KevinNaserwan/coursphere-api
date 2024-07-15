package bookcategory

import (
	"github.com/gin-gonic/gin"
	bookCategoryService "github.com/kevinnaserwan/coursphere-api/internal/service/book_category"
)

type bookCategoryController struct {
	BookCategoryService bookCategoryService.Service
}

func NewBookCategoryController(router *gin.RouterGroup, bookCategoryService bookCategoryService.Service) {
	controller := &bookCategoryController{
		BookCategoryService: bookCategoryService,
	}

	router.POST("/create", controller.Insert)
	router.GET("/all", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.DELETE(":id", controller.Delete)
}
