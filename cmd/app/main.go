package main

import (
	"log"

	"github.com/Gabriel-Yuzo/reservas/docs"
	database "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql"
	models "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/models"
	repository "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/repository"
	"github.com/Gabriel-Yuzo/reservas/internal/config"
	ginHttp "github.com/Gabriel-Yuzo/reservas/internal/http"
	usecase "github.com/Gabriel-Yuzo/reservas/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for a reservations system.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectMySQL(cfg.MysqlURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	db.AutoMigrate(&models.User{})

	userRepo := repository.InitializeRepositories(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	router := ginHttp.NewRouter(userUsecase)
	// Documentação Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor
	if err := router.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
