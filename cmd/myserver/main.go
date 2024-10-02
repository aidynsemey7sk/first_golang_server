package main

import (
	"first_server/pkg/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

// Промедуточный слой для записи логов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Создаем обертку для записи статуса ответа
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		// Логируем запрос и статус ответа
		log.Printf("Запрос: метод=%s, URL=%s, с IP=%s, статус=%d\n", r.Method, r.URL.Path, r.RemoteAddr, rw.statusCode)
		log.Printf("Время выполнения: %s\n", time.Since(start))
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

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

	// Обработка маршрутов
	//http.Handle("/", loggingMiddleware(http.HandlerFunc(handlers.Home)))
	//http.Handle("/about", loggingMiddleware(http.HandlerFunc(handlers.About)))
	//http.Handle("/contact", loggingMiddleware(http.HandlerFunc(handlers.Contact)))

	// Обработка статических файлов из директории "static"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			loggingMiddleware(http.HandlerFunc(handlers.Home)).ServeHTTP(w, r)
		case "/about":
			loggingMiddleware(http.HandlerFunc(handlers.About)).ServeHTTP(w, r)
		case "/contact":
			loggingMiddleware(http.HandlerFunc(handlers.Contact)).ServeHTTP(w, r)
		case "/thanks":
			loggingMiddleware(http.HandlerFunc(handlers.Thanks)).ServeHTTP(w, r)

		default:
			loggingMiddleware(http.HandlerFunc(handlers.NotFoundHandler)).ServeHTTP(w, r)
		}
	})

	// Запуск сервера
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
