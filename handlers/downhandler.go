package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadSampleHandler(w http.ResponseWriter, r *http.Request) {
	// Проверка метода
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Открытие файла
	file, err := os.Open("text.txt")
	if err != nil {
		Status404(w)
		return
	}
	// Закрытие файла преждевременно
	defer file.Close()

	// Указывает браузеру на сохранение файла с именем text.txt
	w.Header().Set("Content-Disposition", "attachment; filename=text.txt")
	// Указывает браузеру на тип содержимого, то есть обычный текст
	w.Header().Set("Content-Type", "text/plain")
	// Получение размера файла
	fileInfo, err := file.Stat()
	if err != nil {
		Status500(w)
		return
	}
	fileSize := fileInfo.Size()
	// Установка размера файла
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))
	// Копирует содержимое file в ответ w
	_, err = io.Copy(w, file)
	if err != nil {
		Status500(w)
		return
	}

}
