// Упражнение 2.3
// перепишите функцию PopCount так,
// чтобы она использовала цикл вместо единого выражения.
// Сравните производительность двух версии.
// (в разделе 11.4 показано, как правильно сравнивать
// производительность различных реализации)

package popcount

// pc[i] - количество единичных битов в i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var s byte
	for i := 0; i < 8; i++ {
		s = pc[byte(x>>(i*8))]
	}
	return int(s)
}