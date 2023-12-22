package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type dbHandler interface {
	// 내부적으로만 사용하기 위해 맨 앞에 소문자 사용
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

var handler dbHandler

// 패키지가 initial 될 때, 한 번만 실행됨
func init() {
	handler = newMemoryHandler()
	// handler = newSqliteHandler()
}

// 외부에서 호출할 때 사용
func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete)
}
