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


#### Функция производительности
Первая версия:
```shell
 go test -bench -benchmem
```

Вторая версия:  
```shell
$ go test -bench=./popcount.go
PASS
ok      github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/exercises/chapter2/2.3   0.002s
```
Затраченное время: 0.002s 