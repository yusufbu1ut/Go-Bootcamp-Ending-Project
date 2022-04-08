package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RegisterMiddlewares(r *gin.Engine) {
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.Use(latencyLogger())
}

func latencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("token", c.GetHeader("Authorization"))

		// before request
		log.Println("This is before request logging")
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
