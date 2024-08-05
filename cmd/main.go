package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	server "test"
	"test/config"
	service2 "test/pkg/service"

	"test/pkg/handlers"
	"test/pkg/repository"
)

func main() {

	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("Error initializing config", err)

	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading env file", err)
	}

	repo, err := repository.NewPostgres(config.PostgresConfig())
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	db := repository.NewRepository(repo)
	service := service2.NewService(db)

	handler := handlers.NewHandler(service)

	fmt.Println("Init success")
	srv := &server.Server{}

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("Failed to run server: ", err)
	}
}
