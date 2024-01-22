package main

import "fmt"

func main() {
	var p *int // 포인터 변수값을 초기화하지 않으면, 기본값은 nil
	if p != nil {
		// p가 nil이 아니라는 이야기는, p가 유효한 메모리 주소를 가진다는 이야기
		fmt.Printf("p는 유효한 메모리 주소를 가지고 있습니다.")
	} else {
		fmt.Printf("p는 nil입니다.")
	}
}
