package main

import (
  "fmt"
  "log"
  "github.com/gin-gonic/gin"
  _ "github.com/gin-gonic/gin/binding"
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

func index (c *gin.Context) {
	c.JSON(200, posts)
}

func save (c *gin.Context) {
  var post Post
  c.BindJSON(&post)
  posts = append(posts, post)
  c.JSON(200, Response {
    Message: "Success",
  })
}

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
  fmt.Println("Running server...")

  router := gin.Default()

	router.GET("/", index)
	router.POST("/", save)
  router.GET("/:id", get)
  router.DELETE("/:id", delete)
  router.PUT("/:id", update)
  
	log.Fatal(router.Run(":6969"))
}