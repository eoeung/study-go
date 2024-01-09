package main

import "fmt"

func main() {
	// [문자열 합치기]
	// + / +=

	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}
