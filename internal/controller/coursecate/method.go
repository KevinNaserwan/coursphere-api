package coursecategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Course Category
// @Description	Create Course Category Endpoint
// @Tags			Course Category
// @Accept			json
// @Produce		json
// @Param			request	body		request.CreateCategoryCourseRequest	true	"Create Course Category Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/course_category/create [post]
func (c *courseCategoryController) Create(ctx *gin.Context) {
	var req request.CreateCategoryCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CourseCategoryService.CreateCourseCategory(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course category created successfully"})
}

// @Summary		Get Course Category By ID
// @Description	Get Course Category By ID Endpoint
// @Tags			Course Category
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Course Category ID"
// @Success		200	{object}	http.Response{value=response.CourseCategoryResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course_category/{id} [get]
func (c *courseCategoryController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.CourseCategoryService.GetCourseCategoryByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"course_category": res})
}

// @Summary		Get All Course Categories
// @Description	Get All Course Categories Endpoint
// @Tags			Course Category
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.Response{value=response.CourseCategoryResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course_category/all [get]
func (c *courseCategoryController) GetAll(ctx *gin.Context) {
	res, err := c.CourseCategoryService.GetAllCourseCategories(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"course_categories": res})
}

// @Summary		Update Course Category
// @Description	Update Course Category Endpoint
// @Tags			Course Category
// @Accept			json
// @Produce		json
// @Param			ID		path		string								true	"Course Category ID"
// @Param			request	body		request.UpdateCategoryCourseRequest	true	"Update Course Category Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/course_category/{id} [put]
func (c *courseCategoryController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.UpdateCategoryCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.CourseCategoryService.UpdateCourseCategory(ctx, ID, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course category updated successfully"})
}

// @Summary		Delete Course Category
// @Description	Delete Course Category Endpoint
// @Tags			Course Category
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Course Category ID"
// @Success		200	{object}	http.Response{value=response.JwtToken}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/course_category/{id} [delete]
func (c *courseCategoryController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.CourseCategoryService.DeleteCourseCategory(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Course category deleted successfully"})
}
