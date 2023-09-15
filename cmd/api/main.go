package main

import (
	"Teeverse/cmd/api/docs"
	"Teeverse/pkg/config"
	di "Teeverse/pkg/di"
	"log"
)

func main() {
	docs.SwaggerInfo.Title = "Teeverse"
	docs.SwaggerInfo.Description = "universal tshirts "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "teeverse.online"
	//docs.SwaggerInfo.Host = "localhost:1243"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
