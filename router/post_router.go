package router

import (
	"my-simple-blog/config"
	"my-simple-blog/handler"
	"my-simple-blog/middleware"
	"my-simple-blog/repository"
	"my-simple-blog/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/articles")

	r.Use(middleware.JWTMiddleware())

	r.POST("/", postHandler.CreateArticle)
	r.GET("/", postHandler.GetArticles)
	r.GET("/:id", postHandler.GetArticle)
	r.GET("/q", postHandler.GetArticleByTitle)
	r.PUT("/:id", postHandler.UpdateArticle)
}
