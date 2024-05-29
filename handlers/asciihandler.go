package handler

import (
	"ascii-web/ascii"
	"net/http"
	"os"
	"strings"
	"text/template"
)

type PageData struct {
	Result string
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	//Проверка метода
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Вытаскиваем текст и баннер
	word := r.FormValue("text")
	format := r.FormValue("options")
	word = strings.ReplaceAll(word, "\r", "")
	ban := ascii.ReadFile(format)

	// Проверки на баннер и введенный текст
	if ban == nil {
		Status404(w)
		return
	}

	if word == "" || !IsAscii(word) {
		Status400(w)
		return
	}

	//Создаем файл, куда будет записываться результат
	file, err := os.Create("text.txt")
	if err != nil {
		Status500(w)
		return
	}
	// Заведомо его закрываем
	defer file.Close()

	// Результат преобразования в ASCII
	totalStr := ascii.Ascii(word, ban)

	// Записываем результат в файл
	err = os.WriteFile("text.txt", []byte(totalStr), 0666)
	if err != nil {
		Status500(w)
		return
	}

	// Чтение файла для вывода содержимого
	result, err := os.ReadFile("text.txt")
	if err != nil {
		Status500(w)
		return
	}
	asciiArt := string(result)

	// Загрузка результата на страницу
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		Status500(w)
		return
	}
	data := PageData{Result: asciiArt}
	if err := tmpl.Execute(w, data); err != nil {
		Status500(w)
		return
	}
}

// Проверяем на ASCII
func IsAscii(word string) bool {
	for _, char := range word {
		if char != 10 && (char < 32 || char > 126) {
			return false
		}
	}
	return true
}
