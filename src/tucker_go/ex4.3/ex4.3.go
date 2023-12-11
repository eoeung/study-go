package main

import "fmt"

func main() {
	var a int = 3 // int64
	var b int     // int 기본값인 0이 됨
	var c = 4     // 타입을 적어주지 않았지만, 값을 선언하여 값의 타입을 따라감
	d := 5

	var e = "Hello"
	f := 3.14 // float은 기본값이 float64

	// var g int32 = 1
	// var k int64 = 20

	// l = g+k // invalid operation: g + k (mismatched types int32 and int64)

	fmt.Println(a, b, c, d, e, f)
}
