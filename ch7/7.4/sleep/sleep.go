// Продолжительность определяет флаг time.Duration с указанным именем,
// значением по умолчанию и строкой использования.
// Возвращаемое значение — это адрес переменной time.Duration,
// в которой хранится значение флага. Флаг принимает значение,
// приемлемое для time.ParseDuration.

package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Ожидание %v... ", *period)
	time.Sleep(*period)
	fmt.Println()
}
