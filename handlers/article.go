package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"net/http"

	"github.com/uacademy/blogpost/api_gateway/models"
	"github.com/uacademy/blogpost/api_gateway/proto-gen/blogpost"
)

// CreateArticle godoc
// @Summary     Create article
// @Description create new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.CreateArticleModel true "Article body"
// @Success     201     {object} models.JSONResult{data=models.Article}
// @Failure     400     {object} models.JSONError
// @Failure     500     {object} models.JSONError
// @Router      /v1/article [post]
func (h Handler) CreateArticle(c *gin.Context) {
	var body models.CreateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	article, err := h.grpcClients.Article.CreateArticle(c.Request.Context(), &blogpost.CreateArticleRequest{
		AuthorId: body.AuthorId,
		Content: &blogpost.Content{
			Title: body.Content.Title,
			Body: body.Content.Body,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "OK",
		Data:    article,
	})
}

// GetArticle godoc
// @Summary     Get article
// @Description get article by ID
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} models.JSONResult{data=models.PackedArticleModel}
// @Failure     404 {object} models.JSONError
// @Router      /v1/article/{id} [get]
func (h Handler) GetArticleById(c *gin.Context) {
	id := c.Param("id")

	article, err := h.grpcClients.Article.GetArticleById(c.Request.Context(), &blogpost.GetArticleByIdRequest{
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
		Data:    article,
	})
}

// ListArticles godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    string false "0"
// @Param       limit  query    string false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONResult{data=[]models.Article}
// @Failure     400    {object} models.JSONError
// @Failure     500    {object} models.JSONError
// @Router      /v1/article [get]
func (h Handler) GetArticleList(c *gin.Context) {
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

	articleList, err := h.grpcClients.Article.GetArticleList(c.Request.Context(), &blogpost.GetArticleListRequest{
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
		Data:    articleList,
	})
}

// UpdateArticle godoc
// @Summary     Update article
// @Description update article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body     models.UpdateArticleModel true "Article body"
// @Success     200     {object} models.JSONResult{data=models.Article}
// @Failure     400     {object} models.JSONError
// @Failure     500     {object} models.JSONError
// @Router      /v1/article [put]
func (h Handler) UpdateArticle(c *gin.Context) {
	var body models.UpdateArticleModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{Error: err.Error()})
		return
	}

	article, err := h.grpcClients.Article.UpdateArticle(c.Request.Context(), &blogpost.UpdateArticleRequest{
		Id: body.Id,
		Content: &blogpost.Content{
			Title: body.Content.Title,
			Body: body.Content.Body,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONError{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    article,
	})
}

// DeleteArticle godoc
// @Summary     Delete article
// @Description delete article by ID
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Article ID"
// @Success     200 {object} models.JSONResult{data=models.Article}
// @Failure     400 {object} models.JSONError
// @Router      /v1/article/{id} [delete]
func (h Handler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := h.grpcClients.Article.DeleteArticle(c.Request.Context(), &blogpost.DeleteArticleRequest{
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
		Data:    article,
	})
}
