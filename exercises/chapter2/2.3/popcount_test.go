package popcount

import (
	p "github.com/unixlinuxgeek/The_Go_Programming_Language_Exercises/ch2/popcount"
	"testing"
)

// Первый тест производительности
func BenchmarkPopCount1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p.PopCount(10000000)
	}
}

// Второй тест производительности (используем цикл for)
func BenchmarkPopCount2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(10000000)
	}
}
