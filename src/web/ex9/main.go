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

func (self *SendComponent) Operator(data string) {
	// send data
	sentData = data
}

// 압축
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

// 암호화
type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

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
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")
	fmt.Println(sentData) // 이상한 문자가 나온다.
	// "Hello World"를 암호화하고 압축했기 때문

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sentData)
	fmt.Println(recvData) // Hello World
}
