package controllers

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Review struct {db *sql.DB}

func (rw *Review) setDatabase(db *sql.DB) {
	rw.db = db
}

func (rw *Review) create(w http.ResponseWriter, r *http.Request) {
	data := models.Review{}
	if err := parseJsonFromBody(r.Body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := data.New(rw.db); err != nil {
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
func (rw *Review) readAll(w http.ResponseWriter, r *http.Request) {
	reviews, err := models.GetAllReviews(rw.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all reviews", zap.Error(err))
		return
	}

	handleJsonResponse(w, reviews)
}
func (rw *Review) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	review := models.Review{Id: id}
	err := review.Get(rw.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific review",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}
	handleJsonResponse(w, review)
}
func (rw *Review) update(w http.ResponseWriter, r *http.Request) {

}
func (rw *Review) delete(w http.ResponseWriter, r *http.Request) {

}