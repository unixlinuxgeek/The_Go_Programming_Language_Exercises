### Упражнения 1.11

Выполните ```fetchall``` с длинным списком аргументов, 
таким как образцы, доступные на сайте alexa.com.
Как ведет себя программа, когда веб-сайт просто не отвечает?
(В разделе 8.9 описан механизм отслеживаний таких ситуаций.)

```shell
go run ./fetchaLL.go https://www.youtube.com https://aliexpress.com https://google.com https://www.amazon.com/ https://reddit.com  https://stackoverflow.com/ https://www.alibaba.com/ https://facebook.com https://reddit.com https://windchill95.rssing.com/chan-33428707/all_p1.html
```

Вывод:
```shell
$ go run ./fetchaLL.go https://www.youtube.com https://aliexpress.com https://google.com https://www.amazon.com/ https://reddit.com  https://stackoverflow.com/ https://www.alibaba.com/ https://facebook.com https://reddit.com https://windchill95.rssing.com/chan-33428707/all_p1.html
0.73s  173874 https://stackoverflow.com/&{0xc00002d9e0}
0.74s  178761 https://www.alibaba.com/&{0xc00022a000}
0.80s  269404 https://www.amazon.com/&{0xc00002dbc0}
0.82s   19457 https://google.com&{0xc0005148a0}
1.01s   68780 https://facebook.com&{0xc00059f5c0}
1.26s  720265 https://www.youtube.com&{0xc0003b6060}
1.34s  477440 https://reddit.com&{0xc00033a3c0}
1.94s  478012 https://reddit.com&{0xc0003b6180}
3.26s     940 https://aliexpress.com&{0xc00033b0e0}
19.63s  225720 https://windchill95.rssing.com/chan-33428707/all_p1.html&{0xc0003b6f00}   // сайт не отвечает
19.63s elapsed
```