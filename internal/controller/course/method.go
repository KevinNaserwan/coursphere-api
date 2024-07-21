package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Course
// @Description	Create Course  Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Param			request	body		request.CreateCourseRequest	true	"Create Course  Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/course/create [post]
func (c *courseController) Create(ctx *gin.Context) {
	var req request.CreateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CourseService.CreateCourse(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course created successfully"})
}

// @Summary		Get Course By ID
// @Description	Get Course By ID Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Course ID"
// @Success		200	{object}	http.Response{value=response.CourseResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course/{id} [get]
func (c *courseController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.CourseService.GetCourseByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"course": res})
}

// @Summary		Get All Courses
// @Description	Get All Courses Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.Response{value=response.CourseResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course [get]
func (c *courseController) GetAll(ctx *gin.Context) {
	res, err := c.CourseService.GetAllCourses(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"courses": res})
}

// @Summary		Update Course
// @Description	Update Course Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Param			ID		path		string						true	"Course ID"
// @Param			request	body		request.UpdateCourseRequest	true	"Update Course Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/course/{id} [put]
func (c *courseController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.UpdateCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CourseService.UpdateCourse(ctx, ID, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// @Summary		Delete Course
// @Description	Delete Course Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Course ID"
// @Success		200	{object}	http.Response{value=response.JwtToken}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course/{id} [delete]
func (c *courseController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.CourseService.DeleteCourse(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// @Summary		Add Course Videos
// @Description	Add Course Videos Endpoint
// @Tags			Course
// @Accept			json
// @Produce		json
// @Param			ID		path		string	true	"Course ID"
// @Param			request	body		request.AddCourseVideosRequest	true	"Add Course Videos Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/course/add-video/{id} [post]
func (c *courseController) AddCourseVideos(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.AddCourseVideosRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CourseService.AddCourseVideo(ctx, ID, req.VideoID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Videos added to course successfully"})
}
