package handler

import (
    "net/http"
    "strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/s/")

    if path == "" {
        http.Error(w, "Short code is required", http.StatusBadRequest)
        return
    }

    originalURL, found := shortenService.Lookup(path)
    if !found {
        http.Error(w, "Short code not found", http.StatusNotFound)
        return
    }

    http.Redirect(w, r, originalURL, http.StatusFound) // 302 Redirect
}
