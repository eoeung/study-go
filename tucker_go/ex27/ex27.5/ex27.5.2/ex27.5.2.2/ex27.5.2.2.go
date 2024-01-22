package main

import "fmt"

// [객체 지향 설계 원칙 SOLID]
// 1) Single Responsibility (단일 책임 원칙)
// 2) Open-Closed (개방-폐쇄 원칙)
// 3) Liskov Substitution (리스코프 치환 원칙)
// 4) Interface Segregation (인터페이스 분리 원칙)
// 5) Dependency Inversion (의존 관계 역전 원칙)

// 5) Dependency Inversion (의존 관계 역전 원칙)
//   - 상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써
//     상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.

//   원칙 2) 추상 모듈은 구체화된 모듈에 의존해서는 안된다.
//           구체화된 모듈은 추상 모듈에 의존해야한다.

// 원칙 2) 예제
// - 메일이 오면 알람이 울린다.

// 2) 올바른 예제
// <<interface>> ◆---- <<interface>>
//    Event             EventListener
//     ◇                     ◇
//     |                      |
//    메일                   알람

// (1) 메일은 Event 인터페이스를 구현하고 있고, 알람은 EventListener라는 인터페이스를 구현한다.
// (2) EventListener는 Event와 관계 맺고 있음 → Event가 EventListener를 소유하고 있음
// (3) 메일이라는 구체화된 객체는 알람이라는 구체화된 객체와 관계를 맺지 않음

// (4) 어떤 구체화된 모듈도 구체화된 모듈에 의존적이지 않고,
// (5) 추상화된 모듈도 구체화된 모듈에 의존적이지 않음

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	Onfire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) { // Event 인터페이스 구현
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.Onfire()
}

type Alarm struct {
}

func (a *Alarm) Onfire() { // EventListener 인터페이스 구현
	fmt.Println("알람! 알람!")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

func main() {
	mail.Register(listener)
	mail.OnRecv()
}
