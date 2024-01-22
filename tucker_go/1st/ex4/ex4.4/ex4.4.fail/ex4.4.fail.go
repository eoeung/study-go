package main

func main() {
	// Go는 최강타입 언어 (타입을 엄격하게 따진다.)
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = b // int // 에러 // cannot use b (variable of type float64) as int value in variable declaration
	d := a * b    // int * float64 // 에러 // invalid operation: a * b (mismatched types int and float64)

	var e int64 = 7 // int64
	f := a * e      // int * int64 // 에러 // invalid operation: a * e (mismatched types int and int64

	var g int = b * 3 // 에러 : float64 * int연산 불가능
}
