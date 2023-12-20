package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	var h string
	fmt.Println(h)

	// var y, z int = 1 // Error : assignment mismatch: 2 variables but 1 value
	var x, y, z int
	x, y, z = 1, 2, 3
	fmt.Printf("%d, %d, %d\n", x, y, z)

	var s = 1                  // 타입 추론
	fmt.Printf("%T, %d", s, s) // int, 1
}
