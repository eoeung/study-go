package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b, "f: ", f)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)  // PrintLine()의 약자
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f) // 서식 문자 작성, 서식에 맞춰 출력할 값 작성
}
