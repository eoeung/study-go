package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos(sessionId string) []*Todo
	AddTodo(sessionId string, name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

// DBHandler 인스턴스를 호출한 쪽에서 사용하다가 종료할 수 있도록 Close() 생성
// → 외부로 책임 전가 → 여기에서는 DBHandler가 종료되는 시점이 언제인지 모르기 때문
func NewDBHandler(filepath string) DBHandler {
	// handler = newMemoryHandler()
	return newSqliteHandler(filepath)
}
