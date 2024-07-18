package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

//	@Summary		Create User
//	@Description	Create User Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UserRegisterRequest	true	"Create User Request"
//	@Success		200		{object}	http.Response{value=response.JwtToken}
//	@Failure		400		{object}	http.Error
//	@Failure		404		{object}	http.Error
//	@Failure		500		{object}	http.Error
//	@Router			/user/create [post]
func (c *userController) Register(ctx *gin.Context) {
	var req request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.Register(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

//	@Summary		Login User
//	@Description	Login User Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UserLoginRequest	true	"Login User Request"
//	@Success		200		{object}	http.Response{value=response.JwtToken}
//	@Failure		400		{object}	http.Error
//	@Failure		404		{object}	http.Error
//	@Failure		500		{object}	http.Error
//	@Router			/user/login [post]
func (c *userController) Login(ctx *gin.Context) {
	var req request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.UserService.Login(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

//	@Summary		Get All Users
//	@Description	Get All Users Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.Response{value=[]response.UserResponse}
//	@Failure		400	{object}	http.Error
//	@Failure		404	{object}	http.Error
//	@Failure		500	{object}	http.Error
//	@Router			/user/all [get]
func (c *userController) GetAll(ctx *gin.Context) {
	users, err := c.UserService.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

//	@Summary		Get User By ID
//	@Description	Get User By ID Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	http.Response{value=response.UserResponse}
//	@Failure		400	{object}	http.Error
//	@Failure		404	{object}	http.Error
//	@Failure		500	{object}	http.Error
//	@Router			/user/{id} [get]
func (c *userController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	user, err := c.UserService.GetByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

//	@Summary		Update User
//	@Description	Update User Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"User ID"
//	@Param			request	body		request.UserUpdateRequest	true	"Update User Request"
//	@Success		200		{object}	http.Response{value=response.UserResponse}
//	@Failure		400		{object}	http.Error
//	@Failure		404		{object}	http.Error
//	@Failure		500		{object}	http.Error
//	@Router			/user/{id} [put]

func (c *userController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.Update(ctx, ID, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

//	@Summary		Delete User
//	@Description	Delete User Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	http.Response{value=response.UserResponse}
//	@Failure		400	{object}	http.Error
//	@Failure		404	{object}	http.Error
//	@Failure		500	{object}	http.Error
//	@Router			/user/{id} [delete]
func (c *userController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.UserService.Delete(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

//	@Summary		Resend Verification Email
//	@Description	Resend Verification Email Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@QueryParam		email	query						string	true	"User Email"
//	@Success		200		{object}	http.Response{value=response.UserResponse}
//	@Failure		400		{object}	http.Error
//	@Failure		404		{object}	http.Error
//	@Failure		500		{object}	http.Error
//	@Router			/user/resend-otp [post]
func (uc *userController) ResendOTPByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	err := uc.UserService.ResendOTPByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP has been resent"})
}

//	@Summary		Verify User OTP
//	@Description	Verify User OTP Endpoint
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.UserOTPVerifyRequest	true	"Verify User OTP Request"
//	@Success		200		{object}	http.Response{value=response.UserResponse}
//	@Failure		400		{object}	http.Error
//	@Failure		404		{object}	http.Error
//	@Failure		500		{object}	http.Error
//	@Router			/user/verify-otp [post]
func (uc *userController) VerifyAuthToken(c *gin.Context) {
	var req request.UserOTPVerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.UserService.VerifyAuthToken(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
}
