package route

import "github.com/gin-gonic/gin"

func ClassesRouter(group *gin.RouterGroup) {
	group.GET("/classes", handler)
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{"Status": "OK"})
}
