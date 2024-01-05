package main

import "fmt"

func main() {
	// 3단 ~ 6단까지는 미출력

	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}

		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}
}
