package popcount

// Побитовый сдвиг влево
func BitShiftLeft(x uint64, y uint) uint64 {
	return x << y
}

// Побитовый сдвиг вправо
func BitShiftRight(x uint64, y uint64) uint64 {
	return x >> y
}
