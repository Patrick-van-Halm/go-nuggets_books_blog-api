package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/authenticator"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", BooksHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/books/{id}", BookByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/reviews", ReviewsHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/reviews/{id}", ReviewByIdHandler).Methods(http.MethodGet)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(authenticator.AuthorizationMiddleware)
	http.ListenAndServe(":8080", r)
}