package main

import (
	"fmt"
	"time"
)

func main() {
	// 1초마다 한 번씩, 1이 증가하는 무한 루프

	i := 1
	for { // → for true {...}와 같음
		time.Sleep(time.Second * 1)
		fmt.Println(i)
		i++
	}
}
