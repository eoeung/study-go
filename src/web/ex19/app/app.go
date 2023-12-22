package app

import (
	"net/http"
	"strconv"
	"web/ex19/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "todo.html", http.StatusTemporaryRedirect)
}

// 모든 todoList → json으로 전송할 예정
func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// todoMap에 저장되어 있는 모든 *Todo(포인터 변수)를 list라는 슬라이스에 저장 후 반환
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

// todo 추가
func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := model.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

// 응답 결과
type Success struct {
	Success bool `json:"success"`
}

// todo 삭제
func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

// todo 완료 여부 확인
func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// 이미 클라이언트 단에서 반대 값으로 넘어옴
	// Ex) 체크 시, 미체크 상태인 false로 넘어옴
	complete := r.FormValue("complete") == "true"

	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func MakeHandler() http.Handler {
	rd = render.New()

	r := mux.NewRouter() // goriila/mux
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	return r
}

// ex17에서 진행했던 todoMap을 이용해서 데이터를 저장하는 방식을 분리하기 위해 Refactoring 진행
// → 데이터를 메모리에 가지고 있는 것이 아닌, 파일 DB로 분리/저장
// → 의존성을 깨뜨리기위한 작업

// SQLite : RDBMS, Query, 무료
