package main

import (
	"first_server/pkg/handlers"
	"first_server/pkg/middlewares"
	"log"
	"net/http"
	"os"
)

func main() {
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatalf("Ошибка при создании директории для логов: %v", err)
	}

	// Открытие файла логов
	logFile, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла логов: %v", err)
	}
	defer logFile.Close()

	// Установка вывода логов в файл
	log.SetOutput(logFile)

	// Логирование запуска сервера
	log.Println("Сервер запущен на порту 8080")

	// Обработка статических файлов из директории "static"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			middlewares.LoggingMiddleware(http.HandlerFunc(handlers.Home)).ServeHTTP(w, r)
		case "/about":
			middlewares.LoggingMiddleware(http.HandlerFunc(handlers.About)).ServeHTTP(w, r)
		case "/contact":
			middlewares.LoggingMiddleware(http.HandlerFunc(handlers.Contact)).ServeHTTP(w, r)
		case "/thanks":
			middlewares.LoggingMiddleware(http.HandlerFunc(handlers.Thanks)).ServeHTTP(w, r)

		default:
			middlewares.LoggingMiddleware(http.HandlerFunc(handlers.NotFoundHandler)).ServeHTTP(w, r)
		}
	})

	// Запуск сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
