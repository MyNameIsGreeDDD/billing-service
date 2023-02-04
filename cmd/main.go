package main

import (
	avito_test_case "avito-test-case"
	"avito-test-case/internal/handler"
	"avito-test-case/internal/repository"
	"avito-test-case/internal/service"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed init configs %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		log.Fatalf("failed init env %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("failed init db %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(avito_test_case.Server)
	if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("run server error %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
