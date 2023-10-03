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

Вывод:
```go
$ go run append.go
0 cap=1 [0]
1 cap=2 [0 1]
2 cap=4 [0 1 2]
3 cap=4 [0 1 2 3]
4 cap=8 [0 1 2 3 4]
5 cap=8 [0 1 2 3 4 5]
6 cap=8 [0 1 2 3 4 5 6]
7 cap=8 [0 1 2 3 4 5 6 7]
8 cap=16        [0 1 2 3 4 5 6 7 8]
9 cap=16        [0 1 2 3 4 5 6 7 8 9]
```
