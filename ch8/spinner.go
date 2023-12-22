package main

import (
	"fmt"
	"time"
)

func main() {
	// Оператор go запускает выполнение вызова функции как независимого
	// параллельного потока управления или горутины в том же адресном пространстве.
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // Медленное вычисление
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
