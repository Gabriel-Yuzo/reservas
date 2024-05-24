package main

import (
	"net/http"

	"github.com/Gabriel-Yuzo/reservas/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Criar o roteador Gin
	r := gin.Default()

	cfg := config.LoadConfig()

	// Definir uma rota
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// Iniciar o servidor
	if err := r.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
