### Упражнения 1.10

Найдите веб-сайт, который содержит большое количество данных.
Исследуйте работу кеширования путем двухкратного запуска ```fetchall```,
и сравнения времени запросов. 
Получите ли вы каждый раз одно и то же содержимое?
Измените ```fetchall``` так, чтобы вывод осуществлялся в файл
и чтобы затем его можно было изучить.

Первый запуск:
```shell
$ go run fetchaLL.go http://google.com http://youtube.com http://github.com
0.51s   19393 http://google.com
0.75s  239178 http://github.com
1.21s  771198 http://youtube.com
1.21s elapsed
```

Второй запуск:
```shell
$ go run fetchaLL.go http://google.com http://youtube.com http://github.com
0.45s   19417 http://google.com
0.70s  239178 http://github.com
1.12s  730549 http://youtube.com
1.12s elapsed
```
