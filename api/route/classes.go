package route

import "github.com/gin-gonic/gin"

func ClassesRouter(group *gin.RouterGroup) {
	group.GET("/classes", handler)
}

type Classes struct {
	Name string
	ID   int
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"classes": []Classes{
			{Name: "1-Introdução_aos_contexto", ID: 51},
        },
	})
}
