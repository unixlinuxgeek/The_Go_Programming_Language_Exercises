// Каждый вызов appendInt должен проверить, имеет ли срез достаточную емкость
// для добавления новых элементов в существующий массив.
// Если дп ф-ия расширяет срез большего объема (в пределах исходного массива).

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := 10000

	appendInt(a, b)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	fmt.Println("Len:", zlen)
	fmt.Println("Cap: ", cap(x))
	if zlen <= cap(x) {
		// Имеется место для роста. Расширяем срез.
		z = x[:zlen]
		fmt.Println("Имеется место для роста. Расширяем срез.")
	} else {
		fmt.Println("Место для роста нет. Выделяем новый массив.")
		fmt.Println("Увеличиваем в два раза для линейной амортизационной сложности.")
		// Место для роста нет. Выделяем новый массив.
		// Увеличиваем в два раза для линейной амортизационной сложности.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // Встроенная функция; см. текст раздела
		fmt.Println("zcap: ", zcap)
		fmt.Println("z: ", z)
	}
	z[len(x)] = y
	return z
}
