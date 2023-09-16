// Упражнение 2.4
// Напишите версию PopCount, которая подсчитывает биты с помощью сдвига аргумента
// по всем 64 позициям, проверяя при каждом сдвиге крайний справа бит.
// Сравните производительность этой версии с выборкой из таблицы.
// ```x&(x)```  сбрасывает крайний справа ненулевой бит x.

package popcount

import "fmt"

// pc[i] - количество единичных битов в i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var s byte
	for i := 0; i < 64; i++ {
		c := pc[x&(x)]
		fmt.Println(c)
		s = pc[byte(x>>(i*8))]
	}

	return int(s)
}
