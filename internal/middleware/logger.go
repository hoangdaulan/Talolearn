package middleware

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the request method and path
		method := c.Request.Method
		path := c.Request.URL.Path
		println("Request:", method, path)

		c.Next()

		// Log the response status code
		statusCode := c.Writer.Status()
		println("Status Code:", statusCode)
	}
}
