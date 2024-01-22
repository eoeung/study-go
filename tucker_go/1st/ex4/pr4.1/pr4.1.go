package main

import "fmt"

func main() {
	a := 3                 // int
	var b = 3.1415         // float64
	c := "hello world"     // string
	d := int32(10)         // int32
	var e float32 = 3.1415 // float32

	fmt.Printf("%T :: %T :: %T :: %T :: %T", a, b, c, d, e)
}
