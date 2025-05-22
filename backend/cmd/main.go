package main

import (
	"cryptodiary/config"
	"cryptodiary/db"
	"log"

	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	cfg := config.DefaultConfig()
	err := supertokens.Init(cfg.SuperTokens)
	if err != nil {
		log.Fatalf("Failed to initialize SuperTokens: %v", err)
	}

	_, err = db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	
}
