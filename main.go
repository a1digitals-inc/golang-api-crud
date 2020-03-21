package main

import (
  db "github.com/asif-ir/golang-api-crud/bootstrap/db"
  router "github.com/asif-ir/golang-api-crud/bootstrap/router"
  seeder "github.com/asif-ir/golang-api-crud/seeds"
)

// Entrypoint
func main() {
  db.InitDb()
  seeder.LoadSeeds(db.DB)
  defer db.DB.Close()
  router.InitRouter()
}
