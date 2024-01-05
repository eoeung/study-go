package main

import "fmt"

func main() {
	// 출력
	// *****
	// ****
	// ***
	// **
	// *

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
