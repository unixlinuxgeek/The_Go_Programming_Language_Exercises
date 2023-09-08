// Server1 - Минимальный "echo" -сервер

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик возвращает компонент пути из URL запроса.
func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "URL.Path: %q\n", req.URL.Path)
}
