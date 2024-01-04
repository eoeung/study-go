package main

import "fmt"

const PI = 3.14 // 타입이 없어서 3.14 숫자로만 동작
const FloatPI float64 = 3.14

func main() {
	// [타입 없는 상수]
	// - 상수 선언시, 타입을 명시하지 않아도 됨

	fmt.Printf("%T\n", PI) // float64

	// var a int = PI * 2 // Error : cannot use PI * 2 (untyped float constant 6.28) as int value in variable declaration (truncated)
	var a int = PI * 200 // 오류가 발생하지 않음
	// → 타입이 정해지지 않은 상수인 float64 * int의 계산인데 오류가 발생하지 않는다???
	// => 변수에 복사될 때 타입이 정해지므로, 연산 결과가 628이기 때문에 오류 발생 X

	// var b int = FloatPI * 100 // Error : cannot use FloatPI * 100 (constant 314 of type float64) as int value in variable declaration
	// → 타입이 정해진 상수인 float64 * int간 연산은 불가능하다.

	fmt.Println(a) // 628
}
