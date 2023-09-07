// FToc выводит результаты преобразования двух температур
// по Фаренгейту в температуру по Цельсию.

package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°С\n", freezingF, fToC(freezingF))
	// "32°F = 0°C"
	fmt.Printf("%g°F = %g°С\n", boilingF, fToC(boilingF))
	// "212°F = 100°C"
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
