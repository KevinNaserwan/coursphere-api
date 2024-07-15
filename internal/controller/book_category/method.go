package bookcategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Book Category
// @Description	Create Book Category Endpoint
// @Tags			Book Category
// @Accept			json
// @Produce		json
// @Param			request	body		request.CreateBookCategoryRequest	true	"Create Book Category Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/book-category/create [post]
func (c *bookCategoryController) Insert(ctx *gin.Context) {
	var req request.CreateBookCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.BookCategoryService.Insert(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// @Summary		Get Book Category By ID
// @Description	Get Book Category By ID Endpoint
// @Tags			Book Category
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Book Category ID"
// @Success		200	{object}	http.Response{value=response.BookCategoryDetailResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/book-category/{ID} [get]
func (c *bookCategoryController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.BookCategoryService.GetByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

// @Summary		Get All Book Categories
// @Description	Get All Book Categories Endpoint
// @Tags			Book Category
// @Accept			json
// @Produce		json
// @Success		200	{object}	http.Response{value=response.BookCategoryResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/book-category/all [get]
func (c *bookCategoryController) GetAll(ctx *gin.Context) {
	bookcategories, err := c.BookCategoryService.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": bookcategories})
}

//	@Summary		Delete Book Category
//	@Description	Delete Book Category Endpoint
//	@Tags			Book Category
//	@Accept			json
//	@Produce		json
//	@Param			ID	path		string	true	"Book Category ID"
//	@Success		200	{object}	http.Response{value=response.BookCategoryDetailResponse}
//	@Failure		400	{object}	http.Error
//	@Failure		404	{object}	http.Error
//	@Failure		500	{object}	http.Error
//	@Router			/book-category/{ID} [delete]

func (c *bookCategoryController) Delete(ctx *gin.Context) {
	ID := ctx.Param("ID")

	if err := c.BookCategoryService.Delete(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book Category deleted successfully"})
}
