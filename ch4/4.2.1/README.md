### Функция append

Встроенная функция append добавляет элементы в срез.

Пример:
```go
package main

import "fmt"

func main() {
	var runes []rune

	for _, r := range "Hello こんにちは" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ' ' 'こ' 'ん' 'に' 'ち' 'は']
}
```

[append.go](https://github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/blob/main/ch4/4.2.1/append.go)