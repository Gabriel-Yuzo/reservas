package main

import (
	"log"

	database "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql"
	models "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/models"
	repository "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/repository"
	"github.com/Gabriel-Yuzo/reservas/internal/config"
	ginHttp "github.com/Gabriel-Yuzo/reservas/internal/http"
	usecase "github.com/Gabriel-Yuzo/reservas/internal/usecase"
)

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
	// docs.SwaggerInfo.BasePath = "/api"
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor
	if err := router.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
