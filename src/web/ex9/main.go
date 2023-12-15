package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

// 전송
type SendComponent struct{}

// SendComponent 구조체가 Component 인터페이스를 구현한 것이 됨
func (self *SendComponent) Operator(data string) {
	// send data
	// fmt.Printf("SendComponent.Operator() ::: sentData ::: %v\n", data)
	sentData = data
}

// 압축
type ZipComponent struct {
	com Component
}

// ZipComponent 구조체가 Component 인터페이스를 구현한 것이 됨
func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("ZipComponent.Operator() ::: zipData ::: %v\n", string(zipData))
	// fmt.Printf("ZipComponent.Operator() ::: zipData ::: %v\n", zipData)
	self.com.Operator(string(zipData))
}

// 암호화
type EncryptComponent struct {
	key string
	com Component
}

func (encrypt *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), encrypt.key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("EncryptComponent.Operator() ::: encryptData ::: %v\n", string(encryptData))
	// fmt.Printf("EncryptComponent.Operator() ::: encryptData ::: %v\n", encryptData)
	encrypt.com.Operator(string(encryptData))
}

// ##################################
// 복호화
type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

// 압축을 푼다.
type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

// data → encrypt → zip → send
func main() {
	// 암호화하는 component가 압축 component를 가지고 있고,
	// 압축 component는 전송 component를 가지고 있음
	sender := &EncryptComponent{ // 구조체 포인터 생성
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
			// Component인터페이스 타입인 com 필드가 있는데, SendComponent는 Component인터페이스를 구현함
		},
	}

	sender.Operator("Hello World")
	fmt.Printf("main() ::: sentData ::: %v\n", sentData) // 이상한 문자가 나온다.
	// "Hello World"를 암호화하고 압축했기 때문

	fmt.Println()
	// ######################################
	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sentData)
	fmt.Printf("main() ::: recvData ::: %v\n", recvData) // Hello World
}
