package controllers

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Series struct {db *sql.DB}

func (s *Series) setDatabase(db *sql.DB) {
	s.db = db
}

func (s *Series) create(w http.ResponseWriter, r *http.Request) {
	data := models.Series{}
	if err := parseJsonFromBody(r.Body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := data.New(s.db); err != nil {
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
func (s *Series) readAll(w http.ResponseWriter, r *http.Request) {
	series, err := models.GetAllSeries(s.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all series", zap.Error(err))
		return
	}

	handleJsonResponse(w, series)
}
func (s *Series) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	series := models.Series{Id: id}
	err := series.Get(s.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific series",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, series)
}
func (s *Series) update(w http.ResponseWriter, r *http.Request) {

}
func (s *Series) delete(w http.ResponseWriter, r *http.Request) {

}