package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/internal/api/repositories/persistence"
)

func ClassesRouter(group *gin.RouterGroup, redisRepo *persistence.RedisRepository) {
	group.GET("/classes", func(c *gin.Context) {
		handler(c, redisRepo)
	})
}
func handler(c *gin.Context, redisRepo *persistence.RedisRepository) {
	json, err := redisRepo.GetJsonConfig(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch classes data"})
		return
	}
	c.JSON(200, json)
}
