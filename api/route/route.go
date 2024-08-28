package route

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/api/middleware"
)

func Setup(gin *gin.Engine) {

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://localhost:8000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	})

	gin.Use(corsMiddleware)

	pb := gin.Group("/")
	pb.Use(middleware.ValidateApiKey())
	pb.GET("/", Ping)

	media := gin.Group("/media")
	MediaRouter(media)
	ClassesRouter(pb)

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
