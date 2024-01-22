package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	// pop 기능
	t, slice := slice[len(slice)-1], slice[:len(slice)-1]

	fmt.Println(t, slice) // 6 [1 2 3 4 5]
}
