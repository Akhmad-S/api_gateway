package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"google.golang.org/grpc"

	"github.com/uacademy/blogpost/api_gateway/clients"
	"github.com/uacademy/blogpost/api_gateway/config"
	"github.com/uacademy/blogpost/api_gateway/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/uacademy/blogpost/api_gateway/handlers"
	"github.com/uacademy/blogpost/api_gateway/storage"
)

func main() {
	cfg := config.Load()

	var stg storage.StorageI


	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// programmatically set swagger info
	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery

	//template GET method
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil{
		panic(err)
	}

	h := handlers.Handler{
		GrpcClients: grpcClients,
	}

	v1 := r.Group("/v1")
	{
		v1.POST("/article", h.CreateArticle)
		v1.GET("/article/:id", h.GetArticleById)
		v1.GET("/article", h.GetArticleList)
		v1.PUT("/article", h.UpdateArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)

		v1.POST("/author", h.CreateAuthor)
		v1.GET("/author/:id", h.GetAuthorById)
		v1.GET("/author", h.GetAuthorList)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
