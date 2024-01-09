package exinit

import "fmt"

var (
	a = c + b // a는 c와 b가 초기화된 이후 초기화됨
	b = f()   // b는 4가됨
	c = f()   // c는 5가됨
	d = 3     // d는 초기화가 끝난 다음 6이 됨
)

func init() {
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Println("d:", d)
}
