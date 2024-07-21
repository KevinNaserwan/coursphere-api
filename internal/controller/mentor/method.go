package mentor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Mentor
// @Description	Create Mentor Endpoint
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			request	body		request.CreateMentorRequest	true	"Create Mentor Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/mentor/create [post]
func (c *mentorController) Create(ctx *gin.Context) {
	var req request.CreateMentorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.MentorService.CreateMentor(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mentor created successfully"})
}

// @Summary		Get Mentor By ID
// @Description	Get Mentor By ID Endpoint
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Mentor ID"
// @Success		200	{object}	http.Response{value=response.MentorResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/mentor/{id} [get]
func (c *mentorController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.MentorService.GetMentorByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mentor": res})
}

// @Summary		Get All Mentors
// @Description	Get All Mentors Endpoint
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.Response{value=response.MentorResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/mentor/all [get]
func (c *mentorController) GetAll(ctx *gin.Context) {
	mentors, err := c.MentorService.GetAllMentors(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"mentors": mentors})
}

// @Summary		Update Mentor
// @Description	Update Mentor Endpoint
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			request	body		request.UpdateMentorRequest	true	"Update Mentor Request"
// @Success		200		{object}	http.Response{value=response.MentorResponse}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/mentor/update [put]
func (c *mentorController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.UpdateMentorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.MentorService.UpdateMentor(ctx, ID, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mentor updated successfully"})
}

// @Summary		Delete Mentor
// @Description	Delete Mentor Endpoint
// @Tags			Mentor
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Mentor ID"
// @Success		200	{object}	http.Response{value=response.MentorResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/mentor/{id} [delete]
func (c *mentorController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.MentorService.DeleteMentor(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Mentor deleted successfully"})
}
