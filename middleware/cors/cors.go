package corsmiddleware

import (
  "fmt"

  "github.com/gin-gonic/gin"
)

// CORS ...
func CORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("CORS middleware")
    c.Next()
  }
}
