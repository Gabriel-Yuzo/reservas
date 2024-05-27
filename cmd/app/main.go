package main

import (
	"fmt"
	"log"
	"net/http"

	database "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql"
	models "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/models"
	repository "github.com/Gabriel-Yuzo/reservas/infra/adpters/mysql/repository"
	"github.com/Gabriel-Yuzo/reservas/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Criar o roteador Gin
	r := gin.Default()

	cfg := config.LoadConfig()

	db, err := database.ConnectMySQL(cfg.MysqlURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	db.AutoMigrate(&models.User{})

	user := &models.User{
		Username: "gabriel",
		Email:    "batatinha@teste.com",
	}

	userRepo := repository.InitializeRepositories(db)

	if err := userRepo.CreateUser(user); err != nil {
		log.Fatalf("Falha ao criar usuário: %v", err)
	}

	// Teste de obtenção de um usuário
	retrievedUser, err := userRepo.GetUserByID(user.ID)
	if err != nil {
		log.Fatalf("Falha ao obter usuário: %v", err)
	}

	fmt.Printf("Usuário obtido: %+v\n", retrievedUser)

	// Definir uma rota
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// Iniciar o servidor
	if err := r.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
