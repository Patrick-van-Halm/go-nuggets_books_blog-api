package main

import "net/http"

func ApiHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Access granted"))
}