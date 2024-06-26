package mysql

import (
	"fmt"
	"log"

	"github.com/Gabriel-Yuzo/reservas/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(cfg config.MySQLConfig) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Database)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Printf(dataSourceName)
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	log.Println("Conexão com o banco de dados bem-sucedida!")
	return db, nil
}

//teste teste workflows
