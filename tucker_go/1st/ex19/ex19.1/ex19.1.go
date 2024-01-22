package main

import "fmt"

type account struct {
	balance int
}

// 일반 함수로 표현
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드로 표현
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	// [메서드]
	// - 리시버
	//   - 메서드에 속하는 타입
	//   - 해당 메서드에서 매개변수처럼 사용
	//   - 모든 로컬 타입 가능
	//     → 로컬 타입 : 해당 패키지 안에서 type 키워드로 선언된 타입
	//     => 패키지 내 선언된 구조체, 별칭 타입

	a := &account{100}             // balance가 100인 account 포인터 변수 생성
	fmt.Printf("%d \n", a.balance) // 40

	withdrawFunc(a, 30)            // 함수 호출
	fmt.Printf("%d \n", a.balance) // 40

	a.withdrawMethod(30) // 메서드 호출
	// 구조체에서 필드가 해당 구조체에 속하듯이, 메서드도 해당 리시버 타입에 속함
	// → .으로 접근 가능
	fmt.Printf("%d \n", a.balance) // 40

}
