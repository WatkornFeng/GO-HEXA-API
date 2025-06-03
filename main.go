package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/WatkornFeng/go-hexa/adapter/config"
	"github.com/WatkornFeng/go-hexa/adapter/handler"
	"github.com/WatkornFeng/go-hexa/adapter/storage/postgres"
	"github.com/WatkornFeng/go-hexa/adapter/storage/postgres/repository"
	"github.com/WatkornFeng/go-hexa/adapter/storage/redis"

	"github.com/WatkornFeng/go-hexa/core/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}
	slog.Info("Starting the application", "app", config.App.Name, "env", config.App.Env)

	// Init Cache(Redis) service
	cache := redis.New(config.Redis)
	defer cache.Close()

	// Init Database(Postgres)
	dbClient := postgres.New(config.DB)
	db := dbClient.DB()
	defer dbClient.Close()

	// Dependency injection
	// User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cache)
	userHandler := handler.NewUserHandler(userService)

	// Init Fiber
	app := fiber.New()
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:userId", userHandler.GetUser)
	app.Post("/users", userHandler.Register)
	app.Patch("/users/:userId", userHandler.UpdateUser)
	app.Delete("/users/:userId", userHandler.DeleteUser)

	// Start server
	listenAddr := fmt.Sprintf(":%s", config.HTTP.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	if err := app.Listen(listenAddr); err != nil {
		slog.Error("Server error", "error", err)
		os.Exit(1)
	}

}
