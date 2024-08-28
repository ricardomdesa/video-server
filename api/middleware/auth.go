package middleware

import "github.com/gin-gonic/gin"

func ValidateApiKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey := ctx.Request.Header.Get("X-API-Key")
        if apiKey != "SeCrEt824" {
            ctx.AbortWithStatusJSON(401, gin.H{"error": "invalid API key"})
            return
        }
        ctx.Next()
	}
}