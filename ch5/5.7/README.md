### Вариативные функции

Вариативная функция это функция с переменным количеством аргументов.
В теле функции типом vals является срез []int

Функция sum может быть вызвана с любым количеством аргументов.

Пример:
```go
func sum(vals ...int) {
	total := 0
	for _, val := range vals {
	    total += val	
    }
	return total
}
```

Параметр ...int ведет себя в теле функции как срез, но не является срезом.
```go
func f(...int) {}
func g([]int) {}

fmt.Printf("%T\n", f) // "func(...int)"
fmt.Printf("%T\n", f) // "func([]int)" 
```


```go
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Стр. %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr) 
}
linenum, name := 12, "count"
errorf(linenum, "не определен %s", name) // Стр. 12: не определен count
```