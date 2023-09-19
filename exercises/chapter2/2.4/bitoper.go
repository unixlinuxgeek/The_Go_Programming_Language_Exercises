package popcount

// Побитовый сдвиг влево
func BitShiftingLeft(x uint64, y uint) uint64 {
	return x << y
}

// Побитовый сдвиг вправо
func BitShiftingRight(x uint64, y uint64) uint64 {
	return x >> y
}
