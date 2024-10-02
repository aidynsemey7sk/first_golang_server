package handlers_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"first_server/pkg/handlers"
)

func TestContactHandler(t *testing.T) {
	// Создаем буфер для перенаправления логов
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil) // Восстановить вывод логов после теста

	// Создаем новый запрос с методом POST и данными формы
	form := "name=John Doe&email=johndoe@example.com&message=Hello!"
	req, err := http.NewRequest("POST", "/contact", bytes.NewBufferString(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Создаем новый записывающий объект для захвата ответа
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	handler := http.HandlerFunc(handlers.Contact)
	handler.ServeHTTP(rr, req)

	// Проверяем код статуса ответа
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("неверный статус: ожидается %v, получено %v", http.StatusSeeOther, status)
	}

	// Проверяем URL перенаправления
	expectedLocation := "/thanks"
	if location := rr.Header().Get("Location"); location != expectedLocation {
		t.Errorf("неверное местоположение перенаправления: ожидается %v, получено %v", expectedLocation, location)
	}

	// Проверяем запись в логах
	expectedLogMessage := "Получено сообщение от John Doe <johndoe@example.com>: Hello!"
	if !bytes.Contains(buf.Bytes(), []byte(expectedLogMessage)) {
		t.Errorf("не найдено сообщение в логах: ожидается %v", expectedLogMessage)
	}
}
