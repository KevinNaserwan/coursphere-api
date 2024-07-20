package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinnaserwan/coursphere-api/internal/http/request"
)

// @Summary		Create Book
// @Description	Create Book Endpoint
// @Tags			Book
// @Accept			json
// @Produce		json
// @Param			request	body		request.CreateBookRequest	true	"Create Book Request"
// @Success		200		{object}	http.Response{value=response.JwtToken}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/book/create [post]
func (c *bookController) Create(ctx *gin.Context) {
	var req request.CreateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := c.BookService.CreateBook(ctx, req); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book created successfully"})
}

// @Summary		Get Book By ID
// @Description	Get Book By ID Endpoint
// @Tags			Book
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Book ID"
// @Success		200	{object}	http.Response{value=response.BookResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/book/{id} [get]
func (c *bookController) GetByID(ctx *gin.Context) {
	ID := ctx.Param("id")

	res, err := c.BookService.GetBookByID(ctx, ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary		Get All Books
// @Description	Get All Books Endpoint
// @Tags			Book
// @Accept			json
// @Produce		json
// @Param			category_name	query	string	false	"Category Name"
// @Success		200	{object}	http.Response{value=response.BookResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/book [get]
func (c *bookController) GetAll(ctx *gin.Context) {
	categoryName := ctx.Query("category_name")
	res, err := c.BookService.GetAllBooks(ctx, categoryName)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary		Update Book
// @Description	Update Book Endpoint
// @Tags			Book
// @Accept			json
// @Produce		json
// @Param			ID		path		string						true	"Book ID"
// @Param			request	body		request.UpdateBookRequest	true	"Update Book Request"
// @Success		200		{object}	http.Response{value=response.BookResponse}
// @Failure		400		{object}	http.Error
// @Failure		404		{object}	http.Error
// @Failure		500		{object}	http.Error
// @Router			/book/{id} [put]
func (c *bookController) Update(ctx *gin.Context) {
	ID := ctx.Param("id")

	var req request.UpdateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.BookService.UpdateBook(ctx, ID, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary		Delete Book
// @Description	Delete Book Endpoint
// @Tags			Book
// @Accept			json
// @Produce		json
// @Param			ID	path		string	true	"Book ID"
// @Success		200	{object}	http.Response{value=response.BookResponse}
// @Failure		400	{object}	http.Error
// @Failure		404	{object}	http.Error
// @Failure		500	{object}	http.Error
// @Router			/book/{id} [delete]
func (c *bookController) Delete(ctx *gin.Context) {
	ID := ctx.Param("id")

	if err := c.BookService.DeleteBook(ctx, ID); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
