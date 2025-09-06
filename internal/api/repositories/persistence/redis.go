package persistence

import (
	context "context"
	"encoding/json"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/domain"
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

func (r *RedisRepository) SaveJson(ctx context.Context, jsonPath, key string) error {
	f, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	err = r.redisClient.Set(ctx, key, f, 0).Err()
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

func (r *RedisRepository) GetCoursesData(ctx context.Context) (domain.Course, error) {
	value, err := r.redisClient.Get(ctx, "courses_json").Result()
	if err != nil {
		return domain.Course{}, err
	}
	var course domain.Course

	if err := json.Unmarshal([]byte(value), &course); err != nil {
		return domain.Course{}, err
	}

	return course, nil
}