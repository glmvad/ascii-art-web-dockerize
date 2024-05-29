package handler

import (
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Проверка URL
	if r.URL.Path != "/" {
		Status404(w)
		return
	}

	// Загрузка главной страницы
	http.ServeFile(w, r, "templates/index.html")
}
