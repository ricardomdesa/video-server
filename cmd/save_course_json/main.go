package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/config"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
)

func main() {
env := config.NewEnv()
	redisConn := config.NewRedis(env, 0)
	defer redisConn.Close()

	get(redisConn)

}

func save(redisConn *redis.Client) {

	persistenceRepo := persistence.NewRedisRepository(redisConn)
	if err := persistenceRepo.SaveJson(context.Background(), "./tempdata.json", "courses_json"); err != nil {
		log.Fatalf("Failed to save classes JSON: %v", err)
		return
	}
	
}

func get(redisConn *redis.Client) {
	persistenceRepo := persistence.NewRedisRepository(redisConn)
	data, err := persistenceRepo.GetCoursesData(context.Background())
	if err != nil {
		log.Fatalf("Failed to get courses JSON: %v", err)
		return
	}
	fmt.Printf("Courses JSON: %v\n", data)

}

