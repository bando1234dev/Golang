package middlewares

import (
	"golang-BE/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware () gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort()
			return
		}

		c.Next()
	}
}