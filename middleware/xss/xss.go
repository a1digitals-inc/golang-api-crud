package xssmiddleware

import (
  "fmt"

  "github.com/gin-gonic/gin"
)

// XSS ...
func XSS() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("XSS middleware")
    c.Next()
  }
}
