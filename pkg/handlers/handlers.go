package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос: метод=%s, URL=%s, с IP=%s\n", r.Method, r.URL.Path, r.RemoteAddr)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Ошибка загрузки шаблона из пути: %s, ошибка: %v", "templates/index.html", err)
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
		log.Printf("Ошибка при выполнении шаблона: %v", err)
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Printf("Запрос: метод=%s, URL=%s, с IP=%s\n", r.Method, r.URL.Path, r.RemoteAddr)

	tmpl := template.Must(template.ParseFiles("templates/about.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем данные из формы
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// Логируем полученные данные (можно также отправлять на email и т.д.)
		log.Printf("Получено сообщение от %s <%s>: %s\n", name, email, message)

		// Перенаправляем пользователя на страницу "Спасибо"
		http.Redirect(w, r, "/thanks", http.StatusSeeOther)
		return
	}

	// Если это GET запрос, просто рендерим форму
	tmpl := template.Must(template.ParseFiles("templates/contact.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("404: Запрос: метод=%s, URL=%s, с IP=%s\n", r.Method, r.URL.Path, r.RemoteAddr)
	tmpl := template.Must(template.ParseFiles("templates/404.html"))
	w.WriteHeader(http.StatusNotFound) // Устанавливаем статус 404
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}

func Thanks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/thanks.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Ошибка при выполнении шаблона", http.StatusInternalServerError)
	}
}
