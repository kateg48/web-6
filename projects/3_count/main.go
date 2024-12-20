package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var c int

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		s := r.Form.Get("count")
		if n, err := strconv.Atoi(s); err != nil {
			w.WriteHeader(http.StatusBadRequest) //400
			w.Write([]byte("это не число"))
			return
		} else {
			c += n
			w.WriteHeader(200)
			return
		}
	}
	if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(c)))
		w.WriteHeader(http.StatusOK) // 200
		return
	}
}

func main() {
	c = 0
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/count", handler)

	// Запускаем веб-сервер на порту 3333
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
