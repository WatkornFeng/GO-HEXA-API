package port

import (
	"context"
	"time"
)

type CacheRepository interface {
	// Set stores the value in the cache
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	// Get retrieves the value from the cache
	Get(ctx context.Context, key string) ([]byte, error)
	// Close cache server
	Close() error
}
