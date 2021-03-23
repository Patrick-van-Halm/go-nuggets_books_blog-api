package main

import (
	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		isServed := false
		auth := strings.Split(authorization, " ")
		if authorization != "" && len(auth) >= 2 && auth[0] == "Bearer" {
			allowedTokens := RequestAuthorizationTokens()
			if contains(allowedTokens, auth[1]) {
				next.ServeHTTP(w, r)
				isServed = true
			}
		}

		if !isServed{
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}



func RequestAuthorizationTokens() []string {
	return []string{
		"Test",
		"Token2",
	}
}