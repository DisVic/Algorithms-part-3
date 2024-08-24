package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	for i := 0; i < len(numbers); i++ {
		_, _ = fmt.Scan(&numbers[i])
	}
	k := 0
	_, _ = fmt.Scan(&k)
	fmt.Println(binarySearch(numbers, k))
	fmt.Println(binarySearch(numbers, k*2))
}

func binarySearch(numbers []int, x int) int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		if numbers[(left+right)/2] >= x {
			right = (left+right)/2 - 1
		} else {
			left = (left+right)/2 + 1
		}
	}
	if numbers[left] >= x {
		return left + 1
	}
	return -1
}
