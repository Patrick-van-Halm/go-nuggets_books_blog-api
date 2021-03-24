package main

import (
	"encoding/json"
	"fmt"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Book"
	review "github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Review"
	"github.com/gorilla/mux"
	"net/http"
)

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(book.GetAll())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func BookByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book := book.GetWithHash(vars["id"])
	if book == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ReviewsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(review.GetAll())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ReviewByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	b, err := json.Marshal(review.GetWithHash(vars["id"]))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}