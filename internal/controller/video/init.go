package video

import (
	"github.com/gin-gonic/gin"
	videoService "github.com/kevinnaserwan/coursphere-api/internal/service/video"
)

type videoController struct {
	VideoService videoService.Service
}

func NewVideoController(router *gin.RouterGroup, videoService videoService.Service) {
	controller := videoController{
		VideoService: videoService,
	}

	router.POST("/create", controller.Create)
	router.GET("", controller.GetAll)
	router.GET("/:id", controller.GetByID)
	router.PUT("/:id", controller.Update)
	router.DELETE("/:id", controller.Delete)
}
