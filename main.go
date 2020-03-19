package main

import (
  "log"
  "github.com/gin-gonic/gin"
  // ginSwagger"github.com/swaggo/gin-swagger"
  // swaggerFiles "github.com/swaggo/files"
)

type Post struct {
  Id string `json:"id"`
  Title string `json:"title"`
}

type Response struct {
  Message string `json:"message"`
}

var posts = []Post {
  {
    Id: "1",
    Title: "Post 1",
  },
  {
    Id: "2",
    Title: "Post 2",
  },
}

// @Summary Get all posts
// @Accept multipart/form-data
// @Produce json
// @Router / [get]
func index (c *gin.Context) {
	c.JSON(200, posts)
}

// @Summary Save post
// @Description Save a new post
// @Produce json
// @Router / [post]
func save (c *gin.Context) {
  var post Post
  c.BindJSON(&post)
  posts = append(posts, post)
  c.JSON(200, Response {
    Message: "Success",
  })
}

// @Summary Get post
// @Description Get a post with given ID
// @Produce json
// @Router /{id} [get]
func get (c *gin.Context) {
  postId := c.Param("id")
  for _, post := range posts {
    if post.Id == postId {
      c.JSON(200, post)
      return
    }
  }
  c.JSON(404, Response {
    Message: "Post not found",
  })
}

// @Summary Delete post
// @Description Delete a post with given ID
// @Produce json
// @Router /{id} [delete]
func delete (c *gin.Context) {
  postId := c.Param("id")
  for i, post := range posts {
    if post.Id == postId {
      posts = append(posts[:i], posts[i+1:]...)
      c.JSON(200, Response {
        Message: "The post with ID " + postId + " has been deleted successfully",
      })
      return
    }
  }
  c.JSON(404, Response {
    Message: "Post not found",
  })
}

// @Summary Update post
// @Description Update a post with given ID
// @Produce json
// @Router /{id} [put]
func update (c *gin.Context) {
  postId := c.Param("id")
  var updatedPost Post
  c.BindJSON(&updatedPost)
  for i, post := range posts {
    if post.Id == postId {
      posts[i].Title = updatedPost.Title
      c.JSON(200, Response {
        Message: "The post with ID " + postId + " has been updated successfully",
      })
      return
    }
  }
  c.JSON(404, Response {
    Message: "Post not found",
  })
}

// @title Posts API
// @version 1.0
// @description This is a sample posts API.
// @contact.name API Support
// @license.name MIT
// @host localhost
// @BasePath /
func main () {
  router := gin.Default()
  router.GET("/", index)
  router.POST("/", save)
  router.GET("/:id", get)
  router.DELETE("/:id", delete)
  router.PUT("/:id", update)
  log.Fatal(router.Run(":6969"))
}