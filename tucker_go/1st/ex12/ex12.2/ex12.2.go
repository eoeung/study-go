package main

import "fmt"

func main() {
	// [배열의 사용 예제]

	var nums [5]int
	fmt.Printf("nums ::: %v\n", nums)

	days := [3]string{"monday", "tuesday", "wednesday"}
	fmt.Printf("days ::: %v\n", days)

	var temps [5]float64 = [5]float64{24.3, 26.7}
	fmt.Printf("temps ::: %v\n", temps)

	var s = [5]int{1: 10, 3: 30} // idx: val → 특정 index값 변경
	fmt.Printf("s ::: %v\n", s)

	x := [...]int{10, 20, 30}             // ...으로 배열 요소 개수 생략 → 초기화되는 요소 개수가 배열 요소 개수가 됨
	fmt.Printf("x ::: %v ::: %T\n", x, x) // [3]int

	y := [...]int{}
	fmt.Printf("y ::: %v ::: %T\n", y, y) // [0]int

	z := []int{}
	fmt.Printf("z ::: %v ::: %T\n", z, z) // []int

}
