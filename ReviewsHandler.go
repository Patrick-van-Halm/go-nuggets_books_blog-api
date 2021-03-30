package main

import (
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func handleReviewsRoutes(r *mux.Router){
	r.HandleFunc("/api/reviews", reviewsHandler).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/reviews/{id}", reviewByIdHandler).Methods(http.MethodGet)
}

func reviewsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		reviewsHandlerGet(w)
		break
	case http.MethodPost:
		reviewsHandlerPost(w, r.Body)
		break
	}
}

func reviewsHandlerGet(w http.ResponseWriter) {
	reviews, err := classes.GetAllReviews(db)
	if err != nil {
		handleHttpError(w, http.StatusInternalServerError, "an error occurred whilst getting all reviews", zap.Error(err))
		return
	}

	handleJsonResponse(w, reviews)
}

func reviewByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	review := classes.Review{Id: id}
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