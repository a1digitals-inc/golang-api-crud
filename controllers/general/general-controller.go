package generalcontroller

import (
	"errors"
	"fmt"
)

// Post ...
type Post struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// Posts ...
var Posts = []Post{
	{
		Id:    "1",
		Title: "Post 1",
	},
	{
		Id:    "2",
		Title: "Post 2",
	},
}

// AddPost ...
func AddPost(post Post) {
	Posts = append(Posts, post)
}

// GetPost ...
func GetPost(postID string) (Post, error) {
	fmt.Println("Given postID ", postID)
	for _, post := range Posts {
		if post.Id == postID {
			return post, nil
		}
	}
	return Post{}, errors.New("Post not found")
}

// DeletePost ...
func DeletePost(postID string) error {
	for i, post := range Posts {
		if post.Id == postID {
			Posts = append(Posts[:i], Posts[i+1:]...)
			return nil
		}
	}
	return errors.New("Post not found")
}

// UpdatePost ...
func UpdatePost(postID string, updatedPost Post) error {
	for i, post := range Posts {
		if post.Id == postID {
			Posts[i].Title = updatedPost.Title
			return nil
		}
	}
	return errors.New("Post not found")
}
