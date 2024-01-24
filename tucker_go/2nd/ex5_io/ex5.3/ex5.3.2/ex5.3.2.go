package main

import "fmt"

func main() {
	// [Scanf()]
	// - 서식에 맞춘 입력
	// func Scanf(format string, a ...any) (n int, err error)

	var a, b int
	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
