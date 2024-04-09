package main

import (
	"log"

	ingestor "github.com/ScMofeoluwa/ingestor/internal"
	"github.com/ScMofeoluwa/ingestor/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Config error: %s", err)
	}

	server := ingestor.NewServer(cfg)
	server.Start()
}
