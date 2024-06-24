package middlewares

import (
	"context"
	"net/http"
)
type contextKey string

// const jwtToken = "token"
const userIDKey contextKey = "userID"
func Auth() func(http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
				ctx := context.WithValue(r.Context(), userIDKey, "helloworld")
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
		})
	}
}