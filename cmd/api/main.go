package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Привет! Это мой первый Go сервер! 🚀")
		if err != nil {
			return
		}
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := fmt.Fprintf(w, `{"message": "Привет из Go!", "status": "ok"}`)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := fmt.Fprintf(w, "Текущее время: %s", time.Now().Format("11-11-2006 15:04:05"))
		if err != nil {
			return
		}
	})

	log.Println("🚀 Сервер запущен на http://localhost:3000")
	log.Println("📝 Доступные endpoints:")
	log.Println("   GET / - главная страница")
	log.Println("   GET /hello - JSON ответ")
	log.Println("   GET /time - показать время")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
