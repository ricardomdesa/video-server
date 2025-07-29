package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/config"
	"github.com/ricardomdesa/videostr/internal/api/middleware"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
)

func Setup(env *config.Env, redis *redis.Client, gin *gin.Engine) {

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	})

	gin.Use(corsMiddleware)

	pb := gin.Group("/")
	pb.Use(middleware.ValidateApiKey(env))
	pb.GET("/ping", Ping)

	media := gin.Group("/media")
	media.Use(middleware.ValidateApiKey(env))

	redisRepo := persistence.NewRedisRepository(redis)

	MediaRouter(media, redisRepo)
	ClassesRouter(pb, redisRepo)

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
