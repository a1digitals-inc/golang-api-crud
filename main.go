package main

import (
  "log"
  "github.com/gin-gonic/gin"
  "github.com/asif-ir/golang-api-crud/routes"
)

// @title Posts API
// @version 1.0
// @description This is a sample posts API.
// @contact.name API Support
// @license.name MIT
// @host localhost
// @BasePath /
func main () {
  router := gin.Default()
  router.GET("/", routes.Index)
  router.POST("/", routes.Save)
  router.GET("/:id", routes.Get)
  router.DELETE("/:id", routes.Delete)
  router.PUT("/:id", routes.Update)
  log.Fatal(router.Run(":6969"))
}