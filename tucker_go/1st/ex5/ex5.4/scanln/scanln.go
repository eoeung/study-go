package main

import "fmt"

func main() {
	// [Scanln()]
	// - 반드시 enter키로 종료해야함

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

	/*
		3 4
		2 3 4

		Hello 4
		0 expected integer

		4 Hello
		1 expected integer
	*/
}
