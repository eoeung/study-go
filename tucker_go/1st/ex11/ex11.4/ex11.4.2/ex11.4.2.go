package main

import "fmt"

func main() {
	// [중첩 for문]
	//   - i x j번 실행

	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
