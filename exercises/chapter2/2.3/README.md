### Упражнение 2.3

Перепишите функцию PopCount так,
чтобы она использовала цикл вместо единого выражения.
Сравните производительность двух версии.
(в разделе 11.4 показано, как правильно сравнивать
производительность различных реализации)

Устанавливаем зависимости:
```shell
go get github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/popcount
```

#### Тесты производительности

```shell
$ go test -bench=.
goarch: amd64
pkg: popcount
cpu: 12th Gen Intel(R) Core(TM) i5-12450H
BenchmarkPopCount1-12           1000000000               0.2458 ns/op
BenchmarkPopCount2-12           163115790                7.497 ns/op
PASS
ok      popcount        2.242s
```

[popcount_test.go](https://github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/blob/main/exercises/chapter2/2.3/popcount_test.go)