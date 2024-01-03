package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scanln(a, b)  // 메모리 주소를 받아야 한다.
	fmt.Println(a, b) // 0 0
}
