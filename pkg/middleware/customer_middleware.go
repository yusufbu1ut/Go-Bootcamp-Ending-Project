package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/hashing"
	jwtHelper "github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/pkg/jwt"
	"net/http"
	"os"
)

func CustomerMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey, os.Getenv("ENV"))
			if decodedClaims != nil {
				//hashed role checking
				roleCheck := hashing.CheckWordHash("customer", decodedClaims.Role) //role is checking
				if roleCheck {
					c.Next()
					c.Abort()
					return
				}
			}

			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
