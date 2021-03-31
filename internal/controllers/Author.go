package controllers

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Author struct {db *sql.DB}

func (a *Author) setDatabase(db *sql.DB) {
	a.db = db
}

func (a *Author) create(w http.ResponseWriter, r *http.Request) {
	data := models.Author{}
	if err := parseJsonFromBody(r.Body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := data.New(a.db); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while creating row",
			zap.Error(err),
			zap.Any(data.TypeName(), &data),
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("Created")); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst writing a response", zap.Error(err))
	}
}
func (a *Author) readAll(w http.ResponseWriter, r *http.Request) {
	authors, err := models.GetAllAuthors(a.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all authors", zap.Error(err))
		return
	}

	handleJsonResponse(w, authors)
}
func (a *Author) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	author := models.Author{Id: id}
	err := author.Get(a.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific author",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, author)
}
func (a *Author) update(w http.ResponseWriter, r *http.Request) {

}
func (a *Author) delete(w http.ResponseWriter, r *http.Request) {

}