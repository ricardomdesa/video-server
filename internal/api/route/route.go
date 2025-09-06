package route

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ricardomdesa/videostr/config"
	"github.com/ricardomdesa/videostr/internal/api/middleware"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
)

func Setup(env *config.Env, redis *redis.Client, gin *gin.Engine, awsSession *session.Session) {

	// Baixa as chaves públicas do Keycloak na inicialização.
	authMidd := middleware.NewAuthMidd(env)
	err := authMidd.FetchPublicKeys()
	if err != nil {
		log.Fatalf("Falha ao buscar as chaves públicas do Keycloak: %v", err)
	}
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:9003"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	})

	gin.Use(corsMiddleware)

	pb := gin.Group("/")
	pb.GET("/ping", Ping)
	pb.Use(authMidd.AuthMiddleware())

	media := gin.Group("/media")
	media.Use(authMidd.AuthMiddleware())

	redisRepo := persistence.NewRedisRepository(redis)
	s3Repo := persistence.NewS3Repository(awsSession, env.S3Bucket)

	MediaRouter(media, s3Repo)
	ClassesRouter(pb, redisRepo)

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
