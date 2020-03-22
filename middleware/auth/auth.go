package authmiddleware

import (
  "fmt"

  "github.com/gin-gonic/gin"
)

// Authenticated ...
func Authenticated() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("Auth middleware")
    c.Next()
  }
}
