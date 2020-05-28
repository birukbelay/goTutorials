package main

import (
	"juliensmith/httprouter"
	"log"
	"net/http"
)

// Middleware without "github.com/julienschmidt/httprouter"

// Middleware without "github.com/julienschmidt/httprouter"
func StdToStdMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff
		next.ServeHTTP(w, r)
	})
}

// Middleware for a standard handler returning a "github.com/julienschmidt/httprouter" Handle
func StdToJulienMiddleware(next http.Handler) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
		next.ServeHTTP(w, r)
	}
}

// Pure "github.com/julienschmidt/httprouter" middleware
func JulienToJulienMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
		next(w, r, ps)
	}
}

func JulienHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
	}
}

func StdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff
	})
}

func main() {
	router := httprouter.New()
	router.POST("/api/user/create", StdToJulienMiddleware(StdHandler()))
	router.GET("/api/user/create", JulienToJulienMiddleware(JulienHandler()))
	log.Fatal(http.ListenAndServe(":8000", router))
}
