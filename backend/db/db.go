package db

import (
	"cryptodiary/config"
	"cryptodiary/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database represents a database connection interface
type Database interface {
	DB() *gorm.DB
}

// PostgresDB implements the Database interface
type PostgresDB struct {
	db *gorm.DB
}

// DB returns the underlying gorm.DB connection
func (p *PostgresDB) DB() *gorm.DB {
	return p.db
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(cfg config.DatabaseConfig) (Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &PostgresDB{db: db}, nil
}

// Migrate performs database migrations
func (p *PostgresDB) Migrate() error {
	// Auto migrate the Entry model
	if err := p.db.AutoMigrate(&models.Entry{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
