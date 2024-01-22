package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // 초기문; 조건문; 후처리
		fmt.Print(i, ", ")
	}

	// fmt.Println(i) // for문 안에 선언된 i는, for문이 종료되면서 메모리에서 제거됨
}
