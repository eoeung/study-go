package main

import "fmt"

// [함수]
// - Go언어는 값 전달만 가능
// - 변수를 전달할 때, 항상 복사로 이루어짐

func Subtract(a, b int) int { // parameter → 입력한 인수값을 복사
	// var a, b int = 3, 6과 같음
	return a - b // 값 반환 → 복사
	// a - b 값이 복사되어 c에 전달되고 나면, 로컬 변수는 삭제된다.
}

func main() {
	c := Subtract(3, 6) // argument
	fmt.Println(c)
}
