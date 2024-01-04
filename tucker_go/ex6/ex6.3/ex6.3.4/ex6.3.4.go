package main

import (
	"fmt"
	"math/big"
)

func main() {
	// math/big 패키지의 Float 객체 사용해서 정밀도 높이기
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d) // 0.1 0.2 0.3 0.3
	fmt.Println(c.Cmp(d))   // 0
}
