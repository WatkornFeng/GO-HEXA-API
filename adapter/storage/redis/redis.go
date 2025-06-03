package redis

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/WatkornFeng/go-hexa/adapter/config"
	"github.com/WatkornFeng/go-hexa/core/port"
	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	client *redis.Client
}

func New(config *config.Redis) port.CacheRepository {

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	// Check Redis connection
	// Use context.Background() for initial connection validation that is not tied to any request.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Simulate a delay longer than the context timeout
	// slog.Info("Simulating slow Redis...")
	// time.Sleep(2 * time.Second)

	_, err := client.Ping(ctx).Result()
	if err != nil {
		slog.Error("Error initializing Redis server", "error", err)
		os.Exit(1)
	}
	slog.Info("Successfully connected to the Redis server")
	return &redisClient{client}
}

func (r *redisClient) Close() error {
	return r.client.Close()
}
func (r *redisClient) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {

	return r.client.Set(ctx, key, value, ttl).Err()
}
func (r *redisClient) Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()
	bytes := []byte(res)
	return bytes, err
}
