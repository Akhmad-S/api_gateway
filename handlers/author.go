package handlers

import (
	"github.com/gin-gonic/gin"
	
	"net/http"
	"strconv"

	"github.com/uacademy/blogpost/api_gateway/models"
	"github.com/uacademy/blogpost/api_gateway/proto-gen/blogpost"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "Author body"
// @Success     201    {object} models.JSONResult{data=models.Author}
// @Failure     400    {object} models.JSONError
// @Failure     500    {object} models.JSONError
// @Router      /v1/author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	author, err := h.grpcClients.Author.CreateAuthor(c.Request.Context(), &blogpost.CreateAuthorRequest{
		Fullname: body.Fullname,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}

// GetAuthor godoc
// @Summary     Get author
// @Description get author by ID
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author ID"
// @Success     200 {object} models.JSONResult{data=models.Author}
// @Failure     404 {object} models.JSONError
// @Router      /v1/author/{id} [get]
func (h Handler) GetAuthorById(c *gin.Context) {
	id := c.Param("id")

	author, err := h.grpcClients.Author.GetAuthorById(c.Request.Context(), &blogpost.GetAuthorByIdRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}

// ListAuthors godoc
// @Summary     List authors
// @Description get authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    string false "0"
// @Param       limit  query    string false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResult{data=[]models.Author}
// @Failure     400    {object} models.JSONError
// @Failure     500    {object} models.JSONError
// @Router      /v1/author [get]
func (h Handler) GetAuthorList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	authorList, err := h.grpcClients.Author.GetAuthorList(c.Request.Context(), &blogpost.GetAuthorListRequest{
		Offset: int32(offset),
		Limit: int32(limit),
		Search: searchStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    authorList,
	})
}

// UpdateAuthor godoc
// @Summary     Update author
// @Description update author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModel true "Author body"
// @Success     200    {object} models.JSONResult{data=models.Author}
// @Failure     400    {object} models.JSONError
// @Failure     404    {object} models.JSONError
// @Router      /v1/author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var body models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	author, err := h.grpcClients.Author.UpdateAuthor(c.Request.Context(), &blogpost.UpdateAuthorRequest{
		Id: body.Id,
		Fullname: body.Fullname,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}

// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete author by ID
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author ID"
// @Success     200 {object} models.JSONResult{data=models.Author}
// @Failure     400 {object} models.JSONError
// @Router      /v1/author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	author, err := h.grpcClients.Author.DeleteAuthor(c.Request.Context(), &blogpost.DeleteAuthorRequest{
		Id: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    author,
	})
}
