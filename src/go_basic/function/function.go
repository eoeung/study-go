package main

import (
	"fmt"
	"time"
)

func main() {
	// 예제로 배우는 GO 프로그래밍
	// 1) Pass by Value
	msg := "Hello"
	println(msg, &msg) // Hello 0xc000069f48
	say(msg)           // Hello 0xc000069f28
	// → 주소 값이 다르게 출력됨 => msg의 값인 "Hello"가 복사되어 함수 say()에 전달됨

	// 2) Pass by Reference
	say2(&msg)         // 0xc000069f48 Hello 0xc000069f28
	println(msg, &msg) // Changed 0xc000069f48
	// → 주소값이 똑같이 출력됨 (0xc000069f48) => 동일한 주소값에서 값만 변경됨

	// 3) 가변 인자 함수
	say3("This", "is", "a", "book") // [4/4]0xc000069f18
	say3("Hi")                      // [1/1]0xc000069ed8

	// 4) 함수 리턴 값
	// (1) 1개 리턴
	total := sum(1, 7, 3, 5, 9)
	println(total)

	// (2) 복수개 리턴
	count, total := sum2(2, 8, 4, 6, 10)
	println(count, total)

	// (3) Named Return Parameter
	count2, total2, duration := sum3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(count2, total2, duration)

	// go by example
}

// 1) Pass by Value
func say(msg string) {
	println(msg, &msg)
}

// 2) Pass by Reference
func say2(msg *string) { // string값이 저장된 포인터인 msg를 인자로 넘긴다.
	println(msg, *msg, &msg)
	// msg = "1234" // cannot use "1234" (untyped string constant) as *string value in assignment
	// → 포인터 타입에 직접적으로 값 할당은 불가능

	*msg = "Changed" // *msg(=msg의 실제 문자열 값)를 사용해서 메모리 주소에 직접 값을 할당
}

// 3) 가변 인자 함수
func say3(msg ...string) {
	println(msg)
	fmt.Printf("%T", msg) // []string

	for index, value := range msg {
		println(index, value)
	}
}

// 4) 함수 리턴 값
// (1) 1개 리턴
func sum(nums ...int) int {
	s := 0
	for _, value := range nums {
		s += value
	}

	return s
}

// (2) 복수개 리턴
func sum2(nums ...int) (int, int) {
	c, s := 0, 0
	for _, value := range nums {
		c++
		s += value
	}

	return c, s
}

// (3) Named Return Parameter
func sum3(nums ...int) (count, total int, duration int64) {
	start := time.Now()
	for _, value := range nums {
		total += value
	}
	count = len(nums)
	duration = time.Since(start).Milliseconds()
	return
}
