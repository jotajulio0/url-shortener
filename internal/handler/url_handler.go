package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"url-shortener/internal/service"

	urlverifier "github.com/davidmytton/url-verifier"
)

// Must have only two main functions: shorten the original url and redirect short code to its respective url
type ShortenReq struct {
	OriginalURL string `json:"url"`
}

type RedirectReq struct {
	shortCode string `json:"short_code"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 16*1024)
	var req ShortenReq

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.OriginalURL == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	originalURL := req.OriginalURL
	verifier := urlverifier.NewVerifier()
	_, err = verifier.Verify(originalURL)

	if err != nil {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	shortCode := service.GenerateShortCode(req.OriginalURL)

	WriteJson(w, http.StatusCreated, map[string]any{"id": 10, "url": originalURL, "code": shortCode})

}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}

func WriteJson(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Printf("Encode Response: %v", err)
	}
}
