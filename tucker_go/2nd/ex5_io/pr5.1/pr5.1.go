package main

import "fmt"

func main() {
	var a = 345
	var b = 3.1415

	fmt.Printf("%05d\n", a)  // 00345
	fmt.Printf("%5.2f\n", b) //  3.14
}
