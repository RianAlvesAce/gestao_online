package middlewares

import (
	"github.com/RianAlvesAce/gestao_online/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearer_schema = "Bearer "

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(bearer_schema):]

		if !jwt.ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}