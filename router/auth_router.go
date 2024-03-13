package router

import (
	"my-simple-blog/config"
	"my-simple-blog/handler"
	"my-simple-blog/repository"
	"my-simple-blog/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
}
