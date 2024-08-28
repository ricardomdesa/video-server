package route

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/repositories"
)

func ClassesRouter(group *gin.RouterGroup) {
	group.GET("/classes", handler)
}

func handler(c *gin.Context) {
	modulos := repositories.GetJsonConfig("assets/media/mod.json")
	c.JSON(200, modulos)
}
