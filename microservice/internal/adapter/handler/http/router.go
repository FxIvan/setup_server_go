package http

import (
	"log/slog"
	"strings"

	"github.com/fxivan/set_up_server/microservice/internal/adapter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	userHandler UserHandler,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", userHandler.RegisterUserHTTP)
			user.POST("/login", userHandler.LoginUserHTTP)
		}
	}

	return &Router{router}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
