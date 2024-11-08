// api/redirect.go
package handler

import (
	"context"
	"net/http"
	"url-shortner/api/utils"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Handler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[1:] // Trim the leading '/'

	if alias == "" {
		http.Error(w, "Alias not provided", http.StatusBadRequest)
		return
	}

	// Get redis Client
	redisClient := utils.GetRedisClient()

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
