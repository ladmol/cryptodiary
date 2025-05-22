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

	database, err := db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Perform database migration
	postgresDB, ok := database.(*db.PostgresDB)
	if !ok {
		log.Fatalf("Failed to cast database to PostgresDB")
	}

	if err := postgresDB.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
