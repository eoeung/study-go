package app

import (
	"net/http"
	"strconv"
	"web/ex20/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

// 응답 결과
type Success struct {
	Success bool `json:"success"`
}

type AppHandler struct {
	// http.Handler를 Embedded 하고 있어서, 마치 http.Handler와 똑같이 사용 가능
	http.Handler // handler http.Handler와 같은 코드 (Has-a)
	db           model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "todo.html", http.StatusTemporaryRedirect)
}

// 모든 todoList → json으로 전송할 예정
func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// todoMap에 저장되어 있는 모든 *Todo(포인터 변수)를 list라는 슬라이스에 저장 후 반환
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

// todo 추가
func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

// todo 삭제
func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

// todo 완료 여부 확인
func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	// 이미 클라이언트 단에서 반대 값으로 넘어옴
	// Ex) 체크 시, 미체크 상태인 false로 넘어옴
	complete := r.FormValue("complete") == "true"

	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{Success: true})
	} else {
		rd.JSON(w, http.StatusOK, Success{Success: false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter() // goriila/mux
	// AppHandler : 핸들러 등록하고, db 호출할 수 있도록 설정
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")

	return a
}

// ex17에서 진행했던 todoMap을 이용해서 데이터를 저장하는 방식을 분리하기 위해 Refactoring 진행
// → 데이터를 메모리에 가지고 있는 것이 아닌, 파일 DB로 분리/저장
// → 의존성을 깨뜨리기위한 작업

// SQLite : RDBMS, Query, 무료
