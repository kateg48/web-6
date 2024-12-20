package main

import (
	"fmt"
	//"io"
	"net/http" // пакет для поддержки HTTP протокола
	//"os"
	//time"
)

// Обработчик HTTP-запросов
func handler(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Hello,%s!", r.URL.Query().Get("name"))
	w.Write([]byte(s))

}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/api/user", handler)

	// Запускаем веб-сервер на порту
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
