// api/redirect.go
package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[1:] // Trim the leading '/'

	if alias == "" {
		http.Error(w, "Alias not provided", http.StatusBadRequest)
		return
	}

	url, err := redisClient.Get(ctx, alias).Result()
	if err == redis.Nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
