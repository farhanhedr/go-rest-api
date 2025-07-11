package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "github.com/gorilla/mux"
)

type Item struct {
  ID          int    `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}

var items []Item
var idCounter int

func getItems(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
  var item Item
  json.NewDecoder(r.Body).Decode(&item)
  idCounter++
  item.ID = idCounter
  items = append(items, item)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(item)
}

func getItem(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if err != nil {
    http.Error(w, "Invalid ID", http.StatusBadRequest)
    return
  }
  for _, item := range items {
    if item.ID == id {
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  http.Error(w, "Item not found", http.StatusNotFound)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    for i, item := range items {
        if item.ID == id {
            json.NewDecoder(r.Body).Decode(&item)
            items[i] = item
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    for i, item := range items {
        if item.ID == id {
            items = append(items[:i], items[i+1:]...)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(items)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Welcome to the Go REST API!")
  })

  r.HandleFunc("/items", getItems).Methods("GET")
  r.HandleFunc("/items", createItem).Methods("POST")
  r.HandleFunc("/items/{id}", getItem).Methods("GET")
  r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
  r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

  fmt.Println("Server running on port 8088")
  err := http.ListenAndServe(":8088", r)
  if err != nil {
    fmt.Println("Error starting server: ", err)
  }

}








