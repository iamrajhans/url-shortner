package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

var (
	urlStore = sync.Map{} // Thread-safe map for storing URLs
)

type shortenRequest struct {
	URL   string `json:"url"`
	Alias string `json:"alias,omitempty"`
}

type shortenResponse struct {
	Alias string `json:"alias"`
	URL   string `json:"url"`
}

// Initialize the random seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

const aliasLength = 6
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateAlias() string {
	b := make([]byte, aliasLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	parsedURL, err := url.ParseRequestURI(req.URL)
	if err != nil || !parsedURL.IsAbs() {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Validate alias if provided
	var validAlias = regexp.MustCompile("^[a-zA-Z0-9_-]{1,}$")
	if req.Alias != "" && !validAlias.MatchString(req.Alias) {
		http.Error(w, "Invalid alias", http.StatusBadRequest)
		return
	}

	// Generate alias if not provided
	alias := req.Alias
	if alias == "" {
		alias = generateAlias()
	}

	if _, exists := urlStore.Load(alias); exists {
		http.Error(w, "Alias already in use", http.StatusConflict)
		return
	}

	// Store the URL with the alias
	urlStore.Store(alias, req.URL)

	resp := shortenResponse{
		Alias: alias,
		URL:   req.URL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	alias := r.URL.Path[1:] // Trim the leading '/'

	if url, exists := urlStore.Load(alias); exists {
		http.Redirect(w, r, url.(string), http.StatusFound)
		return
	}

	http.Error(w, "URL not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	log.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
