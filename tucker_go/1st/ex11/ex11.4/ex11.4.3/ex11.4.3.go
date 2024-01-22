package main

import "fmt"

func main() {
	// [중첩 for문과 break]

	dan := 2
	b := 1
	for {
		for {
			fmt.Printf("%d * %d = %d\n", dan, b, dan*b)
			b++
			if b == 10 {
				break
			}
		}

		b = 1
		dan++
		if dan == 10 {
			break
		}
	}

	fmt.Println("for문이 종료됐습니다.")
}
