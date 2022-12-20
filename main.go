package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()                          // Хранит в себе все URL-ы и зависимости приложения
	fs := http.FileServer(http.Dir("templates/css"))   // handle для нахождения css файлы
	mux.Handle("/css/", http.StripPrefix("/css/", fs)) // удаление префикса "templates" для доступа именно к css либо html файлам
	mux.HandleFunc("/", Home)                          // handle домашней странички
	mux.HandleFunc("/ascii", AsciiDraw)                // handle для создания странички с результатом
	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
	fmt.Println(http.ListenAndServe(":8080", mux))
}
