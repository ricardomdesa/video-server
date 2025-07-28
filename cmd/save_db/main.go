package main

import (
	"context"
	"path/filepath"

	"github.com/ricardomdesa/videostr/config"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
	log "github.com/sirupsen/logrus"
)
func main() {

	log.Info("Starting video saver...")
	env := config.NewEnv()
	
	redisConn := config.NewRedis(env, 0)
	defer redisConn.Close()

	persistenceRepo := persistence.NewRedisRepository(redisConn)

	data, _ := GetMediaData(context.Background(), persistenceRepo, "video:mod1:1-Instalação:index0.ts")
	if data == nil {
		log.Println("DAta here")
		log.Println(data)
		print(data)
	}else{
		log.Println("Data not found for key video:mod1:1-Instalação:index0.ts")
	}
	return

	// saves the media byte data from /assets/media/mod1/1-Instalação/index0.ts to Redis
	byte := []byte(filepath.Join("assets", "media", "mod1", "1-Instalação", "index0.ts"))
	if err := persistenceRepo.SaveMediaData(context.Background(), "video:mod1:1-Instalação:index0.ts", byte); err != nil {
		log.Fatalf("Failed to save media data: %v", err)
	} else {
		log.Info("Media data saved successfully")
	}
	log.Info("Video saver finished successfully")
	if err := redisConn.Close(); err != nil {
		log.Errorf("Failed to close Redis connection: %v", err)
	} else {
		log.Info("Redis connection closed successfully")
	}
	log.Info("Exiting video saver")
}

func GetMediaData(ctx context.Context, redisRepo *persistence.RedisRepository, key string) ([]byte, error) {
	data, err := redisRepo.GetMediaData(ctx, key)
	if err != nil {
		log.Errorf("Failed to get media data for key %s: %v", key, err)
		return nil, err
	}
	log.Infof("Successfully retrieved media data for key %s", key)
	return data, nil
}	