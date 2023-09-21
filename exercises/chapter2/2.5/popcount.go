// Упражнение 2.5
// Выражение ```x&(x-1)```  сбрасывает крайний справа ненулевой бит x.
// Напишите версию PopCount, которая подсчитывает биты с использованием этого факта,
//и оцените ее производительность.

package popcount

// pc[i] - количество единичных битов в i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount возвращает степень заполнения
// (количество установленных битов) значения x.
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])

}

func v(x uint64) int {
	// x&(x-1)
	xX := x & (x - 1)
	return int(
		pc[byte(xX>>(1*8))] +
			pc[byte(xX>>(2*8))] +
			pc[byte(xX>>(3*8))] +
			pc[byte(xX>>(4*8))] +
			pc[byte(xX>>(5*8))] +
			pc[byte(xX>>(6*8))] +
			pc[byte(xX>>(7*8))])
}

func PopCount3(x int) int {
	cnt := 0
	for x != 0 {
		cnt = x & 1
		x >>= 1
	}
	return cnt
}
