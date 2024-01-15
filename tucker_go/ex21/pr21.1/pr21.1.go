package main

import "fmt"

func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func main() {
	fmt.Printf("1부터 5까지 합은 %d입니다.\n", sum(1, 2, 3, 4, 5))
}
