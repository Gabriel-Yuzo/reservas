package main

import (
	"log"

	ginHttp "github.com/Gabriel-Yuzo/reservas/internal/api"
	"github.com/Gabriel-Yuzo/reservas/internal/config"
	"github.com/Gabriel-Yuzo/reservas/internal/domain/models"
	database "github.com/Gabriel-Yuzo/reservas/internal/infra/adpters/mysql"
	"github.com/Gabriel-Yuzo/reservas/internal/repository"
	"github.com/Gabriel-Yuzo/reservas/internal/usecase"
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
	userRepo := repository.NewUserRepositoryImpl(db)
	userUsecase := usecase.NewUserUsecaseImpl(userRepo)
	router := ginHttp.NewRouter(userUsecase)

	// Iniciar o servidor
	if err := router.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
