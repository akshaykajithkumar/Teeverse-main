package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	config "github.com/thnkrn/go-gin-clean-arch/pkg/config"
	di "github.com/thnkrn/go-gin-clean-arch/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, Err := di.InitializeAPI(config)
	if Err != nil {
		log.Fatal("cannot start server: ", Err)
	} else {
		server.Start()
	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}
}
