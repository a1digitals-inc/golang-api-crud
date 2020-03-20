package main

import (
  "log"

  "github.com/gin-gonic/gin"

  authRouter "github.com/asif-ir/golang-api-crud/routes/auth"
  generalRouter "github.com/asif-ir/golang-api-crud/routes/general"
)

// @title Posts API
// @version 1.0
// @description This is a sample posts API.
// @contact.name API Support
// @license.name MIT
// @host localhost
// @BasePath /
func main() {
  router := gin.Default()
  router.GET("/", generalRouter.Index)
  router.POST("/", generalRouter.Save)
  router.GET("/:id", generalRouter.Get)
  router.DELETE("/:id", generalRouter.Delete)
  router.PUT("/:id", generalRouter.Update)

  router.POST("/auth/login", authRouter.Login)
  router.POST("/auth/authenticate", authRouter.Authenticate)
  router.POST("/auth/register", authRouter.Register)

  log.Fatal(router.Run(":6969"))
}
