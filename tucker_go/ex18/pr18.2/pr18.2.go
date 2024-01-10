package main

import "fmt"

func main() {
	// 처음부터 마지막 두 번째 전까지 잘라내는 구문
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceA := slice[:len(slice)-2]
	fmt.Println(sliceA)
}
