package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
// *account 타입에 속함
func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
// account 타입에 속함
func (a2 account) withdrawValue(amount int) {
	a2.balance -= amount
}

// 변경된 값을 반환하는 값 타입 메서드
// account 타입에 속함
func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3
}

func main() {
	// [포인터 메서드 vs 값 타입 메서드]
	// - 리시버를 값 타입과 포인터로 정의할 수 있다.

	// 1) 포인터 메서드
	//   - 포인터가 가리키고 있는 메모리 주소값 복사

	// 2) 값 타입 메서드
	//   - 리시버 타입의 모든 값 복사
	//   - 리시버 타입이 구조체면, 구조체의 모든 데이터가 복사됨

	var mainA *account = &account{100, "Joe", "Park"} // 포인터 변수
	mainA.withdrawPointer(30)                         // 포인터 메서드 호출
	fmt.Println(mainA.balance)                        // 70
	// → a1과 mainA는 같은 메모리 주소, 즉 같은 인스턴스를 가리킴

	mainA.withdrawValue(20) // 값 타입 메서드 호출
	// withdrawValue() 메서드는 account 타입을 리시버로 받는다.
	// → (*mainA).withdrawValue(20)과 같이 값 타입으로 변환해서 호출해야 한다.
	// => Go에서는 자동으로 mainA의 값으로 변환해서 호출해줌
	fmt.Println(mainA.balance) // 70 // 여전히 70 출력
	// → a2와 mainA는 서로 다른 메모리 주소, 즉 서로 다른 인스턴스를 가리킴

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance) // 50

	mainB.withdrawPointer(30)
	// 마찬가지로 withdrawPointer() 메서드는 *account 타입(*account 포인터)를 리시버로 받는다.
	// → (&mainA).withdrawPointer(30)와 같이 주소 연산자를 사용해서 포인터로 변환 후에 호출해야 한다.
	// => Go에서는 자동으로 mainB의 메모리 주소값으로 변환해서 호출해줌
	fmt.Println(mainB.balance) // 20

	// [포인터 메서드와 값 타입 메서드를 만들어야할 때]
	// 1) 포인터 메서드
	//   - 메서드 내부에서 리시버의 값을 변경시킬 수 있음
	//     → 인스턴스 중심

	// 2) 값 타입 메서드
	//   - 호출하는 쪽과 메서드 내부의 값이 별도의 인스턴스로 독립됨
	//   - 메서드 내부에서 리시버의 값을 변경시킬 수 없음
	//     → 값 중심
}
