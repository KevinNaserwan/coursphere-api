package book

import (
	"github.com/gin-gonic/gin"
	bookService "github.com/kevinnaserwan/coursphere-api/internal/service/book"
)

type bookController struct {
	BookService bookService.Service
}

func NewBookController(router *gin.RouterGroup, bookService bookService.Service) {
	controller := &bookController{
		BookService: bookService,
	}

	router.POST("/create", controller.Create)
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}
