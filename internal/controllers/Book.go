package controllers

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/dto"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Book struct {db *sql.DB}

func (b *Book) setDatabase(db *sql.DB) {
	b.db = db
}

func (b *Book) create(w http.ResponseWriter, r *http.Request) {
	bookCreate := dto.CreateBookDTO{}
	if err := parseJsonFromBody(r.Body, &bookCreate); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := bookCreate.Create(b.db); err != nil {
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
func (b *Book) readAll(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks(b.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all books", zap.Error(err))
		return
	}
	data := make([]*dto.BookDTO, 0)
	for _, book := range books {
		d := dto.BookDTO{}
		err := d.Get(book, b.db)
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
func (b *Book) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book := models.Book{Id: id}
	err := book.Get(b.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific book",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	data := dto.BookDTO{}
	err = data.Get(&book, b.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting book DTO",
			zap.Error(err),
			zap.Any("book", book),
		)
		return
	}
	handleJsonResponse(w, data)
}
func (b *Book) update(w http.ResponseWriter, r *http.Request) {

}
func (b *Book) delete(w http.ResponseWriter, r *http.Request) {

}