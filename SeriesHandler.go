package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleSeriesGetRoutes(get *mux.Router) {
	get.HandleFunc("", seriesHandlerGetAll)
	get.HandleFunc("/{id}", seriesGetHandler)
}

func handleSeriesPostRoutes(post *mux.Router) {
	post.HandleFunc("", seriesPostHandler)
}

func seriesPostHandler(w http.ResponseWriter, r *http.Request) {
	handleCreate(&models.Series{}, w, r)
}

func seriesHandlerGetAll(w http.ResponseWriter, _ *http.Request) {
	books, err := models.GetAllSeries(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all series", zap.Error(err))
		return
	}

	handleJsonResponse(w, books)
}

func seriesGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	series := models.Series{Id: id}
	err := series.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific series",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}

	handleJsonResponse(w, series)
}