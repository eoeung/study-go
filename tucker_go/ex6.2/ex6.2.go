package main

import (
	"fmt"
	"math"
)

// a, b(a,b모두 float64 타입)을 인자로 받고, bool타입을 return
func equal(a, b float64) bool {
	fmt.Println(math.Nextafter(a, b)) // 0.3
	return math.Nextafter(a, b) == b  // 1비트 만큼 이동한다.
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", a+b, c, equal(a+b, c))
}
