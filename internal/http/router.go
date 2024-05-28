package http

import (
	"github.com/Gabriel-Yuzo/reservas/internal/http/handler"
	"github.com/Gabriel-Yuzo/reservas/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewRouter(userUsecase usecase.UserUsecase) *gin.Engine {
	router := gin.Default()

	// Definindo as rotas e associando os handlers
	api := router.Group("/api")
	{
		userHandler := handler.NewUserHandler(userUsecase)
		api.POST("/users", userHandler.CreateUser)
		api.GET("/users/:id", userHandler.GetUser)
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	}

	return router
}
