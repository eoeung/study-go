package main

import "fmt"

// 의존 관계 역전 원칙 적용해보기
// 문제

type Player struct {
	name string
}

type Monster struct {
	hp int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Attack(m *Monster) {
	m.DealDamage(p, 100)
}

func (m *Monster) DealDamage(attacker *Player, damage int) {
	m.hp -= damage
	if m.hp < 0 {
		fmt.Println(attacker.Name(), "가 나를 죽였다.")
	}
}
