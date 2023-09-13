### Упражнение 2.2

Напишите программу общего назначения для преобразования единиц,
аналогичную cf, которые считывает числа, из аргументов командной строки,
(или из стандартного ввода, если аргументы командной строки отсутствуют)
и преобразует каждое число в другие едйницы, как температуру - в градусы Цельсия и Фаренгейта,
длину в футы и метры, вес в фунты и килограммы и.т.д.

Устанавливаем зависимости:
```shell
github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/lenconv
github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/tempconv
github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/wtconv
```

Запускаем:
```shell
$ go run ./conv.go 1
1 ℉ = -17.22 ℃ , 1 ℃ = 33.8 ℉ 
1 ft = 0.3 m, 1 m = 3.3 ft
1 lb = 0.45 kg, 1 kg = 2.2 lb
```