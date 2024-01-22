package main

import "fmt"

const Y int = 3

func main() {
	// [배열 요소 읽고 쓰기]

	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
