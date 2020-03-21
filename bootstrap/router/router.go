package bootstraprouter

import (
  "log"

  "github.com/gin-gonic/gin"

  authRouter "github.com/asif-ir/golang-api-crud/routes/auth"
  generalRouter "github.com/asif-ir/golang-api-crud/routes/general"
)

// InitRouter ...
// @title Posts API
// @version 1.0
// @description This is a sample posts API.
// @contact.name API Support
// @license.name MIT
// @host localhost
// @BasePath /
func InitRouter() {
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "UP",
    })
  })
  postsRouterGroup := router.Group("/posts")
  {
    postsRouterGroup.GET("/", generalRouter.Index)
    postsRouterGroup.POST("/", generalRouter.Save)
    postsRouterGroup.GET("/:id", generalRouter.Get)
    postsRouterGroup.DELETE("/:id", generalRouter.Delete)
    postsRouterGroup.PUT("/:id", generalRouter.Update)
  }
  authRouterGroup := router.Group("/auth")
  {
    authRouterGroup.POST("/authenticate", authRouter.Authenticate)
    authRouterGroup.POST("/register", authRouter.Register)
    authRouterGroup.POST("/login", authRouter.Login)
  }
  log.Fatal(router.Run(":6969"))
}
