package postmodel

import (
  "errors"

  db "github.com/asif-ir/golang-api-crud/bootstrap/db"
  "github.com/jinzhu/gorm"
)

// Post ...
type Post struct {
  Id    string `gorm:"size:255;not null;" json:"id"`
  Title string `gorm:"size:255;not null;unique" json:"title"`
}

// Find ...
func Find() ([]Post, error) {
  var err error
  posts := []Post{}
  err = db.DB.Debug().Model(Post{}).Limit(100).Find(&posts).Error
  if err != nil {
    return []Post{}, err
  }
  return posts, nil
}

// FindOne ...
func FindOne(postID string) (Post, error) {
  post := Post{}
  err := db.DB.Debug().Model(Post{}).Where("id = ?", postID).Take(&post).Error
  if err != nil {
    return Post{}, err
  }
  return post, nil
}

// Save ...
func (p *Post) Save() error {
  var err error
  err = db.DB.Debug().Model(&Post{}).Create(&p).Error
  return err
}

// Update ...
func Update(postID string, updatedPost Post) error {
  err := db.DB.Debug().Model(Post{}).Where("id = ?", postID).Updates(Post{Title: updatedPost.Title}).Error
  return err
}

// Delete ...
func Delete(postID string) error {
  var post Post
  err := db.DB.Debug().Model(Post{}).Where("id = ?", postID).Take(&Post{}).Delete(&post).Error
  if err != nil {
    if gorm.IsRecordNotFoundError(err) {
      return errors.New("Post not found")
    }
  }
  return err
}
