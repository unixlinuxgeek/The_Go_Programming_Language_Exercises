// Каждый вызов appendInt должен проверить, имеет ли срез достаточную емкость
// для добавления новых элементов в существующий массив.
// Если дп ф-ия расширяет срез большего объема (в пределах исходного массива).

// Вывод:
// $ go run append.go
// 0 cap=1 [0]
// 1 cap=2 [0 1]
// 2 cap=4 [0 1 2]
// 3 cap=4 [0 1 2 3]
// 4 cap=8 [0 1 2 3 4]
// 5 cap=8 [0 1 2 3 4 5]
// 6 cap=8 [0 1 2 3 4 5 6]
// 7 cap=8 [0 1 2 3 4 5 6 7]
// 8 cap=16        [0 1 2 3 4 5 6 7 8]
// 9 cap=16        [0 1 2 3 4 5 6 7 8 9]

package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// Имеется место для роста. Расширяем срез.
		z = x[:zlen]
	} else {
		// Место для роста нет. Выделяем новый массив.
		// Увеличиваем в два раза для линейной амортизационной сложности.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Встроенная функция; см. текст раздела
	}
	z[len(x)] = y
	return z
}
