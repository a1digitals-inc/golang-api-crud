package main

import (
  "fmt"
  "net/http"
  "log"
  "encoding/json"
  "github.com/gorilla/mux"
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

func index (w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(posts)
}

func save (w http.ResponseWriter, r *http.Request) {
  var msg = Response {
    Message: "Not implemented yet",
  }
  json.NewEncoder(w).Encode(msg)
}

func get (w http.ResponseWriter, r *http.Request) {
  postId := mux.Vars(r)["id"]

  for _, post := range posts {
    if post.Id == postId {
      json.NewEncoder(w).Encode(post)
      return
    }
  }
  var error = Response {
    Message: "Post not found",
  }
  json.NewEncoder(w).Encode(error)
}

func delete (w http.ResponseWriter, r *http.Request) {
  postId := mux.Vars(r)["id"]

  for i, post := range posts {
    if post.Id == postId {
      posts = append(posts[:i], posts[i+1:]...)
      var res = Response {
        Message: "The post with ID " + postId + " has been deleted successfully",
      }
      json.NewEncoder(w).Encode(res)
      return
    }
  }
  var error = Response {
    Message: "Post not found",
  }
  json.NewEncoder(w).Encode(error)
}

func update (w http.ResponseWriter, r *http.Request) {
  var msg = Response {
    Message: "Not implemented yet",
  }
  json.NewEncoder(w).Encode(msg)
}

func main () {
  fmt.Println("Running server")

  router := mux.NewRouter()

  router.HandleFunc("/", index).Methods("GET")
  router.HandleFunc("/", save).Methods("POST")
  router.HandleFunc("/{id}", get).Methods("GET")
  router.HandleFunc("/{id}", delete).Methods("DELETE")
  router.HandleFunc("/{id}", update).Methods("PUT")

  log.Fatal(http.ListenAndServe(":6969", router))
}