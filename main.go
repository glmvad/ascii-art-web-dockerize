package main

import (
 "ascii-web/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/ascii-art", handler.AsciiHandler)
	http.HandleFunc("/download", handler.DownloadSampleHandler)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
