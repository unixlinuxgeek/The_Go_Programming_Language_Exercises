// ### Упражнение 2.2
// Напишите программу общего назначения для преобразования единиц,
// аналогичную cf, которые считывает числа, из аргументов командной строки,
// (или из стандартного ввода, если аргументы командной строки отсутствуют)
// и преобразует каждое число в другие едйницы, как температуру - в градусы Цельсия и Фаренгейта,
// длину в футы и метры, вес в фунты и килограммы и.т.д.

package main

import (
	"fmt"
	"github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/lenconv"
	"github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/tempconv"
	"github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/wtconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, _ := strconv.Atoi(arg)

		cl := tempconv.Cel(t)
		fr := tempconv.Far(t)

		mt := lenconv.Meter(t)
		ft := lenconv.Foot(t)

		// Фунты в килограммы
		kg := wtconv.Kg(t)
		// Килограммы в фунты
		lb := wtconv.Lb(t)

		// температуру - в градусы Цельсия и Фаренгейта
		fmt.Printf("%s = %.2f ℃ , %s = %s \n", fr, tempconv.FToC(fr), cl, tempconv.CToF(cl))

		// длину - в футы и метры
		fmt.Printf("%s = %s, %s = %s\n", ft, lenconv.FToM(ft), mt, lenconv.MToF(mt))

		// вес в фунты и килограммы
		fmt.Printf("%s = %s, %s = %s\n", lb, wtconv.LbToKg(lb), kg, wtconv.KgToLb(kg))
	}
}
