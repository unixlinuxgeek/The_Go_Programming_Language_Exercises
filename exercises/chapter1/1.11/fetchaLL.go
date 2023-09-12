// Упражнения 1.11
// Выполните ```fetchall``` с длинным списком аргументов,
// таким как образцы, доступные на сайте alexa.com.
// Как ведет себя программа, когда веб-сайт просто не отвечает?
// (В разделе 8.9 описан механизм отслеживаний таких ситуаций.)

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan *os.File)
	for _, url := range os.Args[1:] {
		go fetchFile(url, ch) // Запуск go-подпрограммы
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//	func fetch(url string, ch chan<- string) {
//		start := time.Now()
//		resp, err := http.Get(url)
//		if err != nil {
//			ch <- fmt.Sprint(err) // Отправка в канал ch
//			return
//		}
//		nbytes, err := io.Copy(io.Discard, resp.Body)
//		resp.Body.Close() // Исключение утечки ресурсов
//
//		if err != nil {
//			ch <- fmt.Sprintf("while reading %s: %v", url, err)
//			return
//		}
//		secs := time.Since(start).Seconds()
//		ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
//	}

func fetchFile(url string, ch chan<- *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		//ch <- fmt.Sprint(err) // Отправка в канал ch
		ch <- os.NewFile(0, fmt.Sprint(err)) // Отправка в канал ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // Исключение утечки ресурсов

	if err != nil {
		ch <- os.NewFile(0, fmt.Sprintf("while reading %s: %v", url, err))
		//ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %7d %s", secs, nbytes, url)
	f, err := os.OpenFile("out.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	_, err = f.Write([]byte(fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)))
	if err != nil {
		fmt.Printf("Error while reading %s: %v\n", url, err)
	}
	ch <- f
}
