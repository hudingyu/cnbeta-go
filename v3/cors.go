package v3

import "github.com/gin-gonic/gin"

func Cors(fn gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.Request.Header.Get("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		fn(c)
	}
}
