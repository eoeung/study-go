package main

import "fmt"

func main() {
	// [슬라이스 순회]
	// - 역시 배열과 동일

	var slice = []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		slice[i] += 10
	}
	fmt.Println(slice)

	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println(slice)
}
