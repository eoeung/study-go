package main

import "fmt"

const (
	C1 = iota
	C2
	C3
)

const (
	D1 = iota + 1
	D2
	D3
)

func main() {
	fmt.Println(C3, D3) // 2 3
}
