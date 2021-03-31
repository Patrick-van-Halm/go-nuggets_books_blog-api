package controllers

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Genre struct {db *sql.DB}

func (g *Genre) setDatabase(db *sql.DB) {
	g.db = db
}

func (g *Genre) create(w http.ResponseWriter, r *http.Request) {
	data := models.Genre{}
	if err := parseJsonFromBody(r.Body, &data); err != nil {
		handleHttpError(w, http.StatusInternalServerError, "error while parsing json", zap.Error(err))
		return
	}

	if err := data.New(g.db); err != nil {
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
func (g *Genre) readAll(w http.ResponseWriter, r *http.Request) {
	genres, err := models.GetAllGenres(g.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all genres", zap.Error(err))
		return
	}

	handleJsonResponse(w, genres)
}
func (g *Genre) read(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	genre := models.Genre{Id: id}
	err := genre.Get(g.db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific genre",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, genre)
}
func (g *Genre) update(w http.ResponseWriter, r *http.Request) {

}
func (g *Genre) delete(w http.ResponseWriter, r *http.Request) {

}