package middleware

import (
	"net/http"
	"strings"
)

func JWT(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		jwt := strings.Split(header, " ")[1]
		if validate(jwt) {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func validate(wjt string) bool {
	return true
}
