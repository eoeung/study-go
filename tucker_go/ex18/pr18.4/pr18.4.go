package main

import (
	"fmt"
	"sort"
)

// 1) 선수 데이터를 표현하는 구조체를 만든다.
// 2) 높은 득점부터 낮은 순으로 정렬

type Player struct {
	Name    string
	Age     int
	Score   int
	PassAcc float64
}

// []Player는 Len(), Less(), Swap()를 구현하지 않으므로, 새로운 타입을 만들어서 메소드를 정의
type Players []Player

func (p Players) Len() int { return len(p) }

// func (p Players) Less(i, j int) bool { return p[i].Score > p[j].Score } // 높은 득점부터 낮은 순으로 정렬
func (p Players) Less(i, j int) bool { return p[i].Score < p[j].Score } // 높은 득점부터 낮은 순으로 정렬
func (p Players) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	player := []Player{
		{"나통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}

	// sort.Sort()함수는 Len(), Less(), Swap()을 구현해야함
	sort.Sort(sort.Reverse(Players(player)))
	fmt.Println(player)
}
