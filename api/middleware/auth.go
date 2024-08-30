package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardomdesa/videostr/config"
)

func ValidateApiKey(env *config.Env) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Request.Header.Get("X-API-Key")
		if apiKey != env.ApiKey {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid API key"})
			return
		}
		ctx.Next()
	}
}
