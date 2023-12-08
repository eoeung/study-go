package main

import "fmt"

func main() {
	// Go는 최강타입 언어 (타입을 엄격하게 따진다.)
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // float64의 값을 int로 변환
	d := float64(a) * b // int값을 float64로 변환

	var e int64 = 7 // int64
	f := a * int(e) // int64의 값을 int로 변환

	fmt.Println(a, b, c, d, e, f)
}
