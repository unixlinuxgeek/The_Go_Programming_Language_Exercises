package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

// Вывод:
// $ go run sha256.go
// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
// true
// [32]uint8
