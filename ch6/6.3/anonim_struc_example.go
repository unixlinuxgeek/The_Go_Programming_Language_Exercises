// Создание экземпляра анонимной структуры с начальными значениями

package main

import "fmt"

func main() {

	car := struct {
		Model int
		Mark  string
	}{
		Model: 5,
		Mark:  "BMW",
	}

	fmt.Println(car.Mark, car.Model)
}
