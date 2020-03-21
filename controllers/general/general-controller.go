package generalcontroller

import (
  postModel "github.com/asif-ir/golang-api-crud/models/post"
)

// AddPost ...
func AddPost(post postModel.Post) error {
  return post.Save()
}

// GetAllPosts ...
func GetAllPosts() ([]postModel.Post, error) {
  return postModel.Find()
}

// GetPost ...
func GetPost(postID string) (postModel.Post, error) {
  return postModel.FindOne(postID)
}

// DeletePost ...
func DeletePost(postID string) error {
  return postModel.Delete(postID)
}

// UpdatePost ...
func UpdatePost(postID string, updatedPost postModel.Post) error {
  return postModel.Update(postID, updatedPost)
}
