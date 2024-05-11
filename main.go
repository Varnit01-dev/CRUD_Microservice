package main

import (
	"CRUD_Microservices/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/posts", handlers.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
