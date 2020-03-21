package seeder

import (
  "log"

  postModel "github.com/asif-ir/golang-api-crud/models/post"
  "github.com/jinzhu/gorm"
)

var posts = []postModel.Post{
  postModel.Post{
    Title: "Title 1",
    Id:    "a",
  },
  postModel.Post{
    Title: "Title 2",
    Id:    "b",
  },
}

// LoadSeeds ...
func LoadSeeds(db *gorm.DB) {
  var err error
  // Drop tables first
  err = db.Debug().DropTableIfExists(postModel.Post{}).Error
  if err != nil {
    log.Fatalf("cannot drop table: %v", err)
  }
  // Migrate afresh
  err = db.Debug().AutoMigrate(postModel.Post{}).Error
  if err != nil {
    log.Fatalf("Cannot migrate table: %v", err)
  }
  // Insert seed data
  for i := range posts {
    err = db.Debug().Model(postModel.Post{}).Create(&posts[i]).Error
    if err != nil {
      log.Fatalf("Cannot seed posts table: %v", err)
    }
  }
}
