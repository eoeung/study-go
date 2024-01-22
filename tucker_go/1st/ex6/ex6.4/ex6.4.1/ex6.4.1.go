package main

import "fmt"

func main() {
	// [대입 연산자]
	var a int
	a = 10

	// var x int
	// var y int
	// x = y = 10 // Error : y = 10이라는 대입 연산은 결과값이 없기 때문에 x에 할당할 값이 없음

	var x int
	var y int
	x = 10
	y = x

	x, y = 3, 4

	// 복수 대입 연산자
	// → 두 값을 서로 바꾸기
	var k int = 10
	var j int = 200

	k, j = j, k

	// 복합 대입 연산자
	var t = 20
	t = t + 2
	t += 2

	// 증감 연산자
	var p int = 1
	p = p + 1
	p += 1
	p++ // 결과값을 반환하지 않음

	p = p - 1
	p -= 1
	p-- // 결과값을 반환하지 않음

	fmt.Println(a, x, y, k, j, t, p)
}
