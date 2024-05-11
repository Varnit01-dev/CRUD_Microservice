package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"CRUD_Microservices/models"

	"github.com/gorilla/mux"
)

var posts []models.Post

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = len(posts) + 1
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	for _, post := range posts {
		if post.ID == postID {
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	http.NotFound(w, r)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updatedPost models.Post
	_ = json.NewDecoder(r.Body).Decode(&updatedPost)

	for i, post := range posts {
		if post.ID == postID {
			updatedPost.ID = postID
			posts[i] = updatedPost
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}

	http.NotFound(w, r)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	for i, post := range posts {
		if post.ID == postID {
			posts = append(posts[:i], posts[i+1:]...)
			fmt.Fprintf(w, "Post with ID %d is deleted", postID)
			return
		}
	}

	http.NotFound(w, r)
}
