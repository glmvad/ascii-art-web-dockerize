package handler

import (
	"net/http"
	"os"
	"path/filepath"
)

func Status400(w http.ResponseWriter) {
	tmplPath := filepath.Join("templates", "400.html")
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(file)
}

func Status500(w http.ResponseWriter) {
	tmplPath := filepath.Join("templates", "500.html")
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(file)
}

func Status404(w http.ResponseWriter) {
	tmplPath := filepath.Join("templates", "404.html")
	file, err := os.ReadFile(tmplPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write(file)
	
}
