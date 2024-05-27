package http

import (
	"github.com/Gabriel-Yuzo/reservas/internal/http/handler"
	"github.com/Gabriel-Yuzo/reservas/internal/usecase"

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
	}

	return router
}
