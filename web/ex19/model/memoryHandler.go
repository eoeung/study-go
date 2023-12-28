package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
}

// 내부적으로만 사용하기 위해 맨 앞에 소문자 사용
// memoryHandler의 메서드
func (m *memoryHandler) getTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) addTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()} // 새로 만든 Todo의 포인터 변수 생성
	m.todoMap[id] = todo                       // 포인터 변수 저장
	return todo
}

func (m *memoryHandler) removeTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) completeTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}

// 새로운 memoryHandler 인스턴스 만들고, todoMap도 생성해서 포인터 전달
func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}
