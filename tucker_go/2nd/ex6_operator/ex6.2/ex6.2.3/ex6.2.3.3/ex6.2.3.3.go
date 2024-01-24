package main

import (
	"fmt"
	"math/big"
)

func main() {
	// [오차를 무시하는 것이 아닌 정확하게 계산하기]
	// - math/big 패키지 사용

	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))

	// c.Cmp(d)
	// 1) c < d : -1
	// 2) c == d : 0
	// 3) c > d : +1
}
