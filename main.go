package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id,omitempty"`
	name string `json:"name,omitempty"`
}

var items []Item

func GetItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}
func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
	router := mux.NewRouter()
	items = append(items, Item{ID: "1", name: "Item 1"})
	items = append(items, Item{ID: "2", name: "Item 2"})

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
