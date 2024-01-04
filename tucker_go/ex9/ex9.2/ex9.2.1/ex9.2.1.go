package main

import "fmt"

func main() {
	// [if문에 논리 연산자 사용]
	// - &&, ||

	var age = 22

	if age >= 10 && age <= 15 {
		fmt.Println("You are young")
	} else if age > 30 || age < 20 {
		fmt.Println("You are not 20s")
	} else {
		fmt.Println("Best age of your Life")
	}
}
