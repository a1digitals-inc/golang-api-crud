package generalrouter

import (
	"github.com/gin-gonic/gin"

	generalController "github.com/asif-ir/golang-api-crud/controllers/general"
)

// Response ...
// TODO: Move into separate schemas directory
type Response struct {
	Message string `json:"message"`
}

// Index ...
// @Summary Get all posts
// @Accept multipart/form-data
// @Produce json
// @Router / [get]
func Index(c *gin.Context) {
	c.JSON(200, generalController.Posts)
}

// Save ...
// @Summary Save post
// @Description Save a new post
// @Produce json
// @Router / [post]
func Save(c *gin.Context) {
	var post generalController.Post
	c.BindJSON(&post)
	generalController.AddPost(post)
	c.JSON(200, Response{
		Message: "Success",
	})
}

// Get ...
// @Summary Get post
// @Description Get a post with given ID
// @Produce json
// @Router /{id} [get]
func Get(c *gin.Context) {
	postID := c.Param("id")
	post, err := generalController.GetPost(postID)
	if err != nil {
		c.JSON(404, Response{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, post)
}

// Delete ...
// @Summary Delete post
// @Description Delete a post with given ID
// @Produce json
// @Router /{id} [delete]
func Delete(c *gin.Context) {
	postID := c.Param("id")
	err := generalController.DeletePost(postID)
	if err != nil {
		c.JSON(404, Response{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, Response{
		Message: "The post with ID " + postID + " has been deleted successfully",
	})
}

// Update ...
// @Summary Update post
// @Description Update a post with given ID
// @Produce json
// @Router /{id} [put]
func Update(c *gin.Context) {
	postID := c.Param("id")
	var updatedPost generalController.Post
	c.BindJSON(&updatedPost)
	err := generalController.UpdatePost(postID, updatedPost)
	if err != nil {
		c.JSON(404, Response{
			Message: err.Error(),
		})
		return
	}
	c.JSON(200, Response{
		Message: "The post with ID " + postID + " has been updated successfully",
	})
}
