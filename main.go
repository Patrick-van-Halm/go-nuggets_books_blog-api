package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/", ApiHandler)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(AuthorizationMiddleware)
	http.ListenAndServe(":8080", r)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
