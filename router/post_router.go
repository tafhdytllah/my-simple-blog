package router

import (
	"my-simple-blog/config"
	"my-simple-blog/handler"
	"my-simple-blog/repository"
	"my-simple-blog/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/articles")

	r.POST("/", postHandler.Create)
}
