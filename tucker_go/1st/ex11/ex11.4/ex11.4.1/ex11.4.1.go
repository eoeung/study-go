package main

import "fmt"

func main() {
	// [중첩 for문]
	//   - i x j번 실행

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
