package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type contextKey string

// const jwtToken = "token"
const userIDKey contextKey = "userID"
func AddProductDist() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		w.Header().Set("Content-Type", "application/json")
		distID := r.Context().Value(userIDKey).(string)
		fmt.Printf("my value %T and %s",userIDKey,distID)
		response := map[string]string{"message": "okay"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		
	})
}