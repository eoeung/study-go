package main

import "fmt"

func main() {
	str := "Hello\tGo\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str) // 문자열을 기본 서식으로 출력
	/*
		Hello   Go      World
		"Go"is Awesome!
	*/
	fmt.Printf("%s", str) // %s 서식으로 출력
	/*
		Hello   Go      World
		"Go"is Awesome!
	*/
	fmt.Printf("%q", str) // "Hello\tGo\tWorld\n\"Go\"is Awesome!\n" // escape 무시하고 그대로 출력
}
