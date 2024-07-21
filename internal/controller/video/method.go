package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Video
// @Description	Create Video Endpoint
// @Tags			Video
// @Accept			json
// @Produce		json
// @Param			request	body		request.VideoRequest	true	"Create Video Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/video/create [post]
func (c *videoController) Create(ctx *gin.Context) {
	var req request.VideoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.VideoService.CreateVideo(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video created successfully"})
}

// @Summary		Get Video By ID
// @Description	Get Video By ID Endpoint
// @Tags			Video
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Video ID"
// @Success		200	{object}	http.Response{value=response.VideoResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/video/{id} [get]
func (c *videoController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.VideoService.GetVideoByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"video": res})
}

// @Summary		Get All Videos
// @Description	Get All Videos Endpoint
// @Tags			Video
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.Response{value=response.VideoResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/video [get]
func (c *videoController) GetAll(ctx *gin.Context) {
	res, err := c.VideoService.GetAllVideos(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"videos": res})
}

// @Summary		Update Video
// @Description	Update Video Endpoint
// @Tags			Video
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Video ID"
// @Param			request	body		request.VideoRequest	true	"Update Video Request"
// @Success		200	{object}	http.Response{value=response.JwtToken}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/video/{id} [put]
func (c *videoController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.VideoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.VideoService.UpdateVideo(ctx, ID, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video updated successfully"})
}

// @Summary		Delete Video
// @Description	Delete Video Endpoint
// @Tags			Video
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Video ID"
// @Success		200	{object}	http.Response{value=response.JwtToken}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/video/{id} [delete]
func (c *videoController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.VideoService.DeleteVideo(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})
}
