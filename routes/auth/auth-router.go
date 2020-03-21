package authrouter

import (
  "strings"

  "github.com/gin-gonic/gin"

  authController "github.com/asif-ir/golang-api-crud/controllers/auth"
)

var jwtKey = []byte("my_secret_key")

// Response ...
// TODO: Move into separate schemas directory
type Response struct {
  Message string `json:"message"`
}

// Credentials ...
type Credentials struct {
  Password string `json:"password"`
  Username string `json:"username"`
}

// Login ...
// @Summary Login user
// @Produce json
// @Router /auth/login [post]
func Login(c *gin.Context) {
  var creds Credentials
  c.BindJSON(&creds)
  token, err := authController.Login(creds.Username, creds.Password)
  if err != nil {
    c.JSON(500, Response{
      Message: err.Error(),
    })
    return
  }
  c.JSON(200, gin.H{
    "token": token,
  })
}

// Authenticate ...
// @Summary Authenticate user token
// @Produce json
// @Router /auth/authenticate [post]
func Authenticate(c *gin.Context) {
  authorization := c.Request.Header.Get("Authorization")
  if strings.HasPrefix(authorization, "Bearer ") {
    tokenString := strings.Split(authorization, " ")[1]
    token, err := authController.Authenticate(tokenString)
    if token.Valid && err == nil {
      c.JSON(200, gin.H{
        "message": "Success",
      })
      return
    }
  }
  c.JSON(401, gin.H{
    "message": "Token invalid",
  })
}

// Register ...
// @Summary Register user
// @Produce json
// @Router /auth/register [post]
func Register(c *gin.Context) {
  var creds Credentials
  c.BindJSON(&creds)
  err := authController.Register(creds.Username, creds.Password)
  if err != nil {
    c.JSON(401, gin.H{
      "message": "Registration failed",
    })
    return
  }
  c.JSON(200, gin.H{
    "message": "Registration successful",
  })
}
