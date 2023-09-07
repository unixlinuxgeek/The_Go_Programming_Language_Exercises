// Server2 - минимальный "echo" -сервер со счетчиком запросов.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Счетчик, возвращающий количество сделанных запросов.
func counter(rw http.ResponseWriter, res *http.Request) {
	mu.Lock()
	fmt.Fprintf(rw, "Count: %d\n", count)
	mu.Unlock()
}
