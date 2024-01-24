package main

import "fmt"

func main() {
	// [Scanln()]
	// - Scan()과 다른 점은 반드시 'Enter'로 종료해야함

	var a, b int
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
