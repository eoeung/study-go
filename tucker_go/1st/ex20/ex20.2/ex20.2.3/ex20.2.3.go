package main

import (
	"study-go/tucker_go/ex20/ex20.2/fedex"
	"study-go/tucker_go/ex20/ex20.2/koreaPost"
)

// Sender 인터페이스 생성
// Send() 메서드만 있으면 누구나 Sender 인터페이스가 됨
type Sender interface {
	Send(parcel string)
}

// Sender 인터페이스를 입력으로 받는다.
// → 기존에는 Fedex나 우체국 등 한 택배 회사만 입력으로 받았음
func SendBook(name string, sender Sender) {
	sender.Send(name) // Sender 인터페이스는 Send()를 포함하고 있음
}

// 1) SendBook()함수 입장에서는 저 Sender가 Fedex인지 우체국인지 관심이 없음
// 2) 메서드 내부 구현을 알 수도, 알 필요도 없음 → 구현 여부가 중요

func main() {
	// 우체국 전송 객체
	koreaPostSender := &koreaPost.PostSender{}
	SendBook("어린 왕자", koreaPostSender)

	// Fedex 전송 객체
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)

	// 1) 우체국, Fedex 객체 모두 Send()메서드를 구현하고 있으므로 Sender 인터페이스가 될 수 있다.
	// 2) SendBook() 함수는 Sender 인터페이스를 입력으로 받기 때문에 아무런 에러가 발생하지 않는다.
}
