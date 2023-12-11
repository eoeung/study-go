package main

import (
	"fmt"
	"math"
	"reflect"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	// 숫자 상수는 타입이 주어지기 전까지 타입을 가지지 않음
	const n = 500000000
	fmt.Println(reflect.TypeOf(n)) // int

	const d = 3e20 / n
	fmt.Println(d, reflect.TypeOf(d)) // 6e+11 float64

	const k, l = int(d), int64(d)
	fmt.Printf("%d %T\n", k, k) // 600000000000 int

	fmt.Printf("%v %T\n", reflect.TypeOf(k), l) // int int64

	// k = 1 // cannot assign to k (neither addressable nor a map index expression)

	fmt.Println(math.Sin(n))

	const h, g = 1e2, 3e2
	fmt.Println(h, g) // 100

	// 상수 관련, 다른 타입으로 보이는 연산 예제 코드
	a := 40
	f := a / 100.0 // 언뜻보면 int, float 연산이라 불가능해야하지만 에러가 발생하지 않음
	// → 40은 int, 100.0은 상수로 선언되어 타입을 가지지 않는 상태
	// => 정수간 나눗셈으로 간주
	fmt.Println(f, reflect.TypeOf(f))  // 0 int
	fmt.Println(reflect.TypeOf(100.0)) // float64
}
