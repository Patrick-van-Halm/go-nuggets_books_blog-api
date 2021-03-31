package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleReviewsGetRoutes(get *mux.Router) {
	get.HandleFunc("", reviewsGetAllHandler)
	get.HandleFunc("/{id}", reviewsGetHandler)
}

func handleReviewsPostRoutes(post *mux.Router) {
	post.HandleFunc("", reviewsPostHandler)
}

func reviewsPostHandler(w http.ResponseWriter, r *http.Request) {
	handleCreate(&models.Review{}, w, r)
}

func reviewsGetAllHandler(w http.ResponseWriter, _ *http.Request) {
	reviews, err := models.GetAllReviews(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all reviews", zap.Error(err))
		return
	}

	handleJsonResponse(w, reviews)
}

func reviewsGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	review := models.Review{Id: id}
	err := review.Get(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting specific review",
			zap.Error(err),
			zap.String("id", id),
		)
		return
	}
	handleJsonResponse(w, review)
}