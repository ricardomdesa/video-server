package persistence

import (
	context "context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/domain"
	"os"
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
	err := r.redisClient.Set(ctx, key, data, 0).Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *RedisRepository) SaveClassesJson(ctx context.Context, jsonPath string) error {
	f, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	err = r.redisClient.Set(ctx, "classes_json", f, 0).Err()
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

func (r *RedisRepository) GetJsonConfig(ctx context.Context) (domain.Classes, error) {
	value, err := r.redisClient.Get(ctx, "classes_json").Result()
	mod := domain.Classes{}
	if err != nil {
		return domain.Classes{}, err
	}
	if err = json.Unmarshal([]byte(value), &mod); err != nil {
		return domain.Classes{}, err
	}
	return mod, nil
}
