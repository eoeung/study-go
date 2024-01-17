package main

import "fmt"

// error 인터페이스
type error interface {
	Error() string
}

type PasswordError struct { // 에러 구조체 선언
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // Error() 메서드
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8} // Error() 메서드를 구현했으므로 error 인터페이스로 사용될 수 있음
	}

	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { // 인터페이스 변환
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입됐습니다.")
	}
}
