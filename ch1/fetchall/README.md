### Fetchall

Fetchall выполняет параллельную выборку URL и сообщает
о затраченном времени и размере ответа для каждого из них.

Вывод:
```shell
$ go run ./fetchall.go http://google.com
0.44s   19419 http://google.com
0.44s elapsed
```