package main

import (
	"study-go/tucker_go/ex20/ex20.2/fedex"
	"study-go/tucker_go/ex20/ex20.2/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체를 생성
	sender := &koreaPost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
	// Error : cannot use sender (variable of type *koreaPost.PostSender) as *fedex.FedexSender value in argument to SendBook

	// 택배 회사가 많아지거나, 코드가 많아질수록 수정해야하는 부분이 많아짐
	// → 인터페이스 사용으로 해결
}
