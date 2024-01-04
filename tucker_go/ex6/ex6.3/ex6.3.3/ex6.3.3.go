package main

import (
	"fmt"
	"math"
)

// ＊math패키지의 Nextafter() 함수 사용
func equal(a, b float64) bool {
	fmt.Printf("math.Nextafter(%0.18f, %0.18f) ::: %f\n", a, b, math.Nextafter(a, b))
	return math.Nextafter(a, b) == b
}

func main() {
	// [오차를 없애는 더 좋은 방법]
	// - 1비트 차이만큼 비교하기
	// → 가장 작은 차이는, 가장 오른쪽 비트값 1개 만큼이다.
	var a, b, c float64 = 0.1, 0.2, 0.3
	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))

	fmt.Println()

	// math.Nextafter(x, y float64)
	// → x와 y 사이에서, x와 가장 가까운 부동소수점을 찾음
	ttt := math.Nextafter(0.12, 10.2)
	fmt.Printf("##### %0.18f\n", ttt)
	ttt = math.Nextafter(10.2, 0.12)
	fmt.Printf("##### %0.18f\n", ttt)
}
