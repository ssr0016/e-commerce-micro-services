package main

import (
	"log"

	"github.com/ssr0016/ecommmerse-app/config"
	"github.com/ssr0016/ecommmerse-app/internal/api"
)

func main() {

	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}
