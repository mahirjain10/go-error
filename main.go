package main

import (
	"log"
	"net/http"

	"github.com/mahirjain_10/handlers"
	"github.com/mahirjain_10/middlewares"
)

func main() {
	r := http.NewServeMux()
	authMiddleware := middlewares.Auth()

	// Create a custom handler to enforce the POST method
	postOnly := func(handler http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			handler.ServeHTTP(w, r)
		}
	}

	r.Handle("/add-product-distributor", postOnly(authMiddleware(http.HandlerFunc(handlers.AddProductDist()))))

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}