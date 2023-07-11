package app

import (
	"github.com/orbit1987/TodoList/internal/api/http/v1"
	"github.com/orbit1987/TodoList/internal/repository"
	"github.com/orbit1987/TodoList/internal/server"
	"github.com/orbit1987/TodoList/internal/service"
	"github.com/spf13/viper"
	"log"
)

const port = "port"

func Run() {
	if err := initConfig(); err != nil {
		log.Fatalf("viper readInConfig err: %s", err.Error())
	}

	var port = viper.GetString(port)

	var rep = repository.NewRepository()
	var services = service.NewService(rep)
	var handlers = handler.NewHandler(services)
	var router = handlers.InitRouter()

	srv := server.Server{}
	if err := srv.Run(port, router); err != nil {
		log.Fatalf("http service error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}
