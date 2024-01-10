package main

import "fmt"

func main() {
	// Ex) 1부터 15까지 추가한 슬라이스
	var slice []int

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)
}
