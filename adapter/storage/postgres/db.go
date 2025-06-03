package postgres

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/WatkornFeng/go-hexa/adapter/config"
	"github.com/WatkornFeng/go-hexa/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresClient struct {
	db *gorm.DB
}

// New initializes the database connection and performs migration
func New(cfg *config.DB) *postgresClient {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get DB instance", "error", err)
		os.Exit(1)
	}
	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		slog.Error("Database ping failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully connected to the database")

	// Auto migrate
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}

	slog.Info("Successfully migrated the database")
	return &postgresClient{db}
}

func (p *postgresClient) DB() *gorm.DB {
	return p.db
}

func (p *postgresClient) Close() error {

	sqlDB, err := p.db.DB()
	if err != nil {
		slog.Warn("Database didn't close cleanly", "error", err)
		return err
	}
	return sqlDB.Close()
}
