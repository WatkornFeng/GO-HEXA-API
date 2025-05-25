package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/WatkornFeng/go-hexa/adapter/config"
	"github.com/WatkornFeng/go-hexa/adapter/handler"
	"github.com/WatkornFeng/go-hexa/adapter/storage/repository"
	"github.com/WatkornFeng/go-hexa/core/domain"
	"github.com/WatkornFeng/go-hexa/core/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Name)
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
	defer sqlDB.Close()
	slog.Info("Successfully connected to the database", "db", config.DB.Connection)

	// Migrate database
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		slog.Error("Error migrating database", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully migrated the database")

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	// Init Fiber
	app := fiber.New()
	app.Get("/users", userHandler.GetUsers)

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	port := config.HTTP.Port
	if port == "" {
		port = "3000"
	}
	if err := app.Listen(":" + port); err != nil {
		slog.Error("Shutting down due to error", "error", err)
		os.Exit(1) // Means exit with error
	}
}
