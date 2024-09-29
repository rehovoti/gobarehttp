package main

import (
  "net/http"
  "fmt"
  "encoding/json"
)

type User struct {
  Name string `json:"name"`
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("GET /", HandleRoot)
  mux.HandleFunc("POST /user", createUser)
  
  http.ListenAndServe("localhost:8080", mux)
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Kuku")
}

func createUser(w http.ResponseWriter, r *http.Request) {
  var user User
  err := json.NewDecoder(r.Body).Decode(&user) 
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return 
  }
  fmt.Fprint(w,"created %v", user)
}
