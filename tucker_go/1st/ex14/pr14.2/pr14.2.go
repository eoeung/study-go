package main

import "fmt"

type Actor struct {
	Name  string
	HP    int
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	// 탈출 분석이용해서 함수가 종료돼도 메모리가 살아있게 된다.
	actor := &Actor{name, hp, speed} // 초기화한 구조체의 메모리 주소
	return actor
}

func main() {
	var actor = NewActor("금토끼", 99, 100)
	fmt.Println(actor.Speed)
	fmt.Println(actor.Name)
}
