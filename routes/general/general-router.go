package routes

import (
  "github.com/gin-gonic/gin"
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
func Index (c *gin.Context) {
	c.JSON(200, posts)
}

// @Summary Save post
// @Description Save a new post
// @Produce json
// @Router / [post]
func Save (c *gin.Context) {
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
func Get (c *gin.Context) {
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
func Delete (c *gin.Context) {
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
func Update (c *gin.Context) {
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
