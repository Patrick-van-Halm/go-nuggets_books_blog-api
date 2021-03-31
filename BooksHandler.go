package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/dto"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleBooksGetRoutes(get *mux.Router) {
	get.HandleFunc("", booksHandlerGetAll)
	get.HandleFunc("/{id}", bookGetHandler)
}

func handleBooksPostRoutes(post *mux.Router) {
	post.HandleFunc("", booksPostHandler)
}

func booksPostHandler(w http.ResponseWriter, r *http.Request) {
	bookCreate := dto.CreateBookDTO{}
	if err := parseJsonFromBody(r.Body, &bookCreate); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := bookCreate.Create(db); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while creating row",
			zap.Error(err),
			zap.Any("book", &bookCreate),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Created")); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}

func booksHandlerGetAll(w http.ResponseWriter, _ *http.Request) {
	books, err := models.GetAllBooks(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all books", zap.Error(err))
		return
	}
	data := make([]*dto.BookDTO, 0)
	for _, b := range books {
		d := dto.BookDTO{}
		err := d.Get(b, db)
		if err != nil {
			handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting book DTO",
				zap.Error(err),
				zap.Any("book", b),
			)
			return
		}
		data = append(data, &d)
	}

	handleJsonResponse(w, data)
}

func bookGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book := models.Book{Id: id}
	err := book.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific book",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	data := dto.BookDTO{}
	err = data.Get(&book, db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting book DTO",
			zap.Error(err),
			zap.Any("book", book),
		)
		return
	}
	handleJsonResponse(w, data)
}