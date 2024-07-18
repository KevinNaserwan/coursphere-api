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

	router.POST("/book", controller.Create)
	router.GET("/book", controller.GetAll)
	router.GET("/book/:id", controller.GetByID)
	router.PUT("/book/:id", controller.Update)
	router.DELETE("/book/:id", controller.Delete)
}
