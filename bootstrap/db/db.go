package bootstrapdb

import (
  "github.com/jinzhu/gorm"
  // Load sqlite
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB ...
var DB *gorm.DB

// InitDb ...
func InitDb() {
  DB, _ = gorm.Open("sqlite3", "sqlite.db")
}
