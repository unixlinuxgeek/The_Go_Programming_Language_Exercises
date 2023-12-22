// Clock1 является TCP-сервером, периодически выводящем время.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleCon(conn) // Обработка единственного подключения
	}
}

func handleCon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // Например, отключение клиента.
		}
		time.Sleep(1 * time.Second)
	}
}
