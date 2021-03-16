package middleware

import (
	"log"
	"net/http"

	"github.com/sjaureguio/golang-api/clase-3/authorization"
)

type HttpMiddleware func(http.ResponseWriter, *http.Request)

// Log .
func Log(f HttpMiddleware) HttpMiddleware {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("petición: %q, método: %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

// Authentication
func Authentication(f HttpMiddleware) HttpMiddleware {
	return func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden(rw, r)
			return
		}

		f(rw, r)
	}
}

// forbidden .
func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}
