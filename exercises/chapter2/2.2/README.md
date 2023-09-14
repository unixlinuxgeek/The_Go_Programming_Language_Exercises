### Упражнение 2.2

Напишите программу общего назначения для преобразования единиц,
аналогичную cf, которые считывает числа, из аргументов командной строки,
(или из стандартного ввода, если аргументы командной строки отсутствуют)
и преобразует каждое число в другие едйницы, как температуру - в градусы Цельсия и Фаренгейта,
длину в футы и метры, вес в фунты и килограммы и.т.д.

```shell
go mod init github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/exercises/chapter2/2.2
```

Устанавливаем зависимости:
```shell
go get github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/lenconv
go get github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/tempconv
go get github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/wtconv
go get github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/exercises/chapter2/2.2/conv
```

#### Вариант 1
Запускаем:
```shell
$ go run ./conv.go 1
1 ℉ = -17.22 ℃ , 1 ℃ = 33.8 ℉ 
1 ft = 0.3 m, 1 m = 3.3 ft
1 lb = 0.45 kg, 1 kg = 2.2 lb
```

#### Вариант 2
Или можно скомпилировать в бинарник:
```shell
go build ./conv.go
```

Запускаем скомпилированный бинарник:
```shell
./conv.go 1
2 ℉ = -16.67 ℃ , 2 ℃ = 35.6 ℉ 
2 ft = 0.61 m, 2 m = 6.6 ft
2 lb = 0.91 kg, 2 kg = 4.4 lb
```
