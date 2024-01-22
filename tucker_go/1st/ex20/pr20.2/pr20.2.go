package main

import (
	"fmt"
	AwesomeDB "study-go/tucker_go/ex20/pr20.2/awesomeDB"
)

// OurDB 구조체의 모든 공개된 메서드를 이용하는 인터페이스 만들어보기
type DB interface {
	GetData() string
	WriteData(data string)
	Close() error
}

func main() {
	var ourDB DB = &AwesomeDB.OurDB{}
	var yourDB DB = &AwesomeDB.OurDB{}

	fmt.Println(ourDB.GetData())
	fmt.Println(yourDB.GetData())

}
