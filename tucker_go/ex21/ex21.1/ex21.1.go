package main

import "fmt"

func sum(nums ...int) int { // 인수 타입인 int 앞에 ...을 붙여 가변 인수 타입을 선언
	sum := 0

	fmt.Printf("nums 타입 : %T\n", nums) // []int → int 슬라이스 타입
	// → 가변 인수는 함수 내부에서, 해당 타입의 슬라이스로 처리됨
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	// [가변 인수 함수]
	// - 함수 인수 개수가 고정적이지 않은 함수

	fmt.Println()
	fmt.Println(1)
	fmt.Println(1, 2, 3, 4, 5)
	fmt.Println()

	// - ... 키워드 사용
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())

	// fmt.Println()은 여러 타입 인수를 섞어쓸 수 있다.
	fmt.Println(2, "hello", 3.14)
	// → 빈 인터페이스 interface{}
	// => 모든 타입이 빈 인터페이스를 포함하고 있음 → ...interface{}

	fmt.Println()
	Print("hello", true, 3.14, 2)
}

// Print() 함수 대충 구현해보기
func Print(args ...interface{}) { // 모든 타입을 받는 가변 인수
	fmt.Printf("### %v\n", args) // [hello true 3.14 2]

	for n, arg := range args { // var arg interface{} ## var args []interface{}
		fmt.Println(n)
		switch arg.(type) { // 인수의 타입에 따른 동작
		case string:
			fmt.Println(arg)
		case bool:
			fmt.Println(arg)
		case float64:
			fmt.Println(arg)
		case int:
			fmt.Println(arg)
		}
	}
}
