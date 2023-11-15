// Аварийная ситуация: присваивание записи в пустом отображении
// $ go run ./urlValues.go
// en
// 1
// [1 2]
// panic: assignment to entry in nil map
// goroutine 1 [running]:
// net/url.Values.Add(...)
// /usr/local/go/src/net/url/url.go:905
// main.main()
// /home/geek/GoApps/The_Go_Programming_Language_Exercises/ch6/6.2.1/urlValues.go:20 +0x4b7
// exit status 2

package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // Непосредственное создание
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // en
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1" (первое значение)
	fmt.Println(m["item"])     // [1 2] (непосредственное обращение)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // Аварийная ситуация: присваивание записи в пустом отображении
}
