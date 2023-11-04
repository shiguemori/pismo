package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"os"
	"pismo/models"
	"sync"
)

var once sync.Once

type Infra struct {
	DB *gorm.DB
}

func (i *Infra) InitDB() {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", dbHost, dbUser, dbPass, dbName, dbPort)
	i.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Error connecting to database", err)
	}

	slog.Info("Connected to database.")
}

func (i *Infra) SetupTables() {
	err := i.DB.AutoMigrate(&models.Account{})
	if err != nil {
		log.Fatalf("Erro na migração da tabela accounts: %v", err)
	}

	err = i.DB.AutoMigrate(&models.OperationType{})
	if err != nil {
		log.Fatalf("Erro na migração da tabela operation_types: %v", err)
	}

	err = i.DB.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatalf("Erro na migração da tabela transactions: %v", err)
	}
}

func NewInfra() *Infra {
	i := &Infra{}
	once.Do(func() {
		i.InitDB()
		i.SetupTables()
	})
	return i
}
