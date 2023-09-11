### Упражнения 1.10

Найдите веб-сайт, который содержит большое количество данных.
Исследуйте работу кеширования путем двухкратного запуска ```fetchall```,
и сравнения времени запросов. 
Получите ли вы каждый раз одно и то же содержимое?
Измените ```fetchall``` так, чтобы вывод осуществлялся в файл
и чтобы затем его можно было изучить.

Первый запуск:
```shell
$ go run ./fetchaLL.go https://youtube.com https://youtube.com
1.92s  723732 https://youtube.com&{0xc00002c420}
1.92s elapsed
```

Второй запуск:
```shell
$ go run ./fetchaLL.go https://youtube.com https://youtube.com
0.98s  788946 https://youtube.com&{0xc000284000}
1.69s  721021 https://youtube.com&{0xc000284060}
1.69s elapsed
```