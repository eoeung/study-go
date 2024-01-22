package main

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 2) Open-Closed
//   - 확장에는 열려있고, 변경에는 닫혀있다.
//   - 상호 결합도를 줄여서 새 기능을 추가할 때, 기존 구현을 변경하지 않아도 됨
//   → 프로그램에 기능을 추가할 때 기존 코드의 변경을 최소화해야 한다.

// Ex) 개방-폐쇄 원칙을 위반한 예시

func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case Email:
		// 이메일 전송
	case Fax:
		// 팩스 전송
	case PDF:
		// pdf 전송
	case Printer:
		// 프린팅
	}
}

// ＊우편 전송(Post)이 추가되는 경우
// 1) SendReport() 함수를 수정해야함
// → 기존 코드를 변경해야함

// + 단일 책임 원칙에도 위배됨
//   → 전송 이라는 기능에 더해서 여러 방식의 전송을 하고 있으므로 단일 책임이 아님
