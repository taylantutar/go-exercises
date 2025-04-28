package handler

import (
    "encoding/json"
    "net/http"
    "7-url-shortener/internal/service"
)

type ShortenRequest struct {
    URL string `json:"url"`
}

type ShortenResponse struct {
    ShortURL string `json:"short_url"`
}

var shortenService = service.NewShortenService()

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var req ShortenRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if req.URL == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    shortCode := shortenService.Shorten(req.URL)

    resp := ShortenResponse{
        ShortURL: "http://localhost:8080/s/" + shortCode,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
