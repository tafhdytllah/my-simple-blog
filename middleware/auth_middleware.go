package middleware

import (
	"my-simple-blog/errorhandler"
	"my-simple-blog/helper"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if tokenStr == "" {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenStr)
		if err != nil {
			errorhandler.HandleError(c, &errorhandler.UnauthorizedError{
				Message: err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}
