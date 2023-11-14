package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	max := 0
	if len(vals) > 0 {
		for i, val := range vals {
			if i == 0 {
				max = val
			}
			if val > max {
				max = val
			}
		}
	}
	return max
}

func min(vals ...int) int {
	min := 0
	if len(vals) > 0 {
		for i, val := range vals {
			if i == 0 {
				min = val
			}
			if val < min {
				min = val
			}
		}
	}
	return min
}

func main() {
	min := min(1, 2, 3)
	fmt.Println("min:", min) // min: 1
	max := max(1, 2, 3)
	fmt.Println("max:", max) // max: 3
}
