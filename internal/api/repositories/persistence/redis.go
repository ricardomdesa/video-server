package persistence

import (
	"github.com/go-redis/redis/v8"
	context "context"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisRepository {
	return &RedisRepository{
		redisClient: redisClient,
	}
}

// Saves media data to Redis key value store
func (r *RedisRepository) SaveMediaData(ctx context.Context, key string, data interface{}) error {
	// Implementation for saving media data
	err := r.redisClient.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *RedisRepository) GetMediaData(ctx context.Context, key string) ([]byte, error) {
	val, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	return val, nil
}