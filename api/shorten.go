// api/shorten.go
package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"url-shortner/api/utils"
)

var ctx = context.Background()

type shortenRequest struct {
	URL   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

type shortenResponse struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Generate alias if not provided
	alias := req.Alias
	if alias == "" {
		alias = generateAlias()
	}

	// Get redis Client
	redisClient := utils.GetRedisClient()
	// Check if alias already exists
	exists, err := redisClient.Exists(ctx, alias).Result()
	if err != nil {
		http.Error(w, "Internal server error redisClient"+err.Error(), http.StatusInternalServerError)
		return
	}

	if exists == 1 {
		http.Error(w, "Alias already in use", http.StatusConflict)
		return
	}

	// Store the URL with the alias
	err = redisClient.Set(ctx, alias, req.URL, 0).Err()
	if err != nil {
		http.Error(w, "Failed to store URL", http.StatusInternalServerError)
		return
	}

	resp := shortenResponse{
		Alias: alias,
		URL:   req.URL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

const aliasLength = 6
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateAlias() string {
	b := make([]byte, aliasLength)
	for i := range b {
		b[i] = letters[randInt(len(letters))]
	}
	return string(b)
}

func randInt(n int) int {
	// Use time-based seed for simplicity
	return int(time.Now().UnixNano() % int64(n))
}
