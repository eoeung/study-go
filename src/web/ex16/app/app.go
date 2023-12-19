package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "todo.html", http.StatusTemporaryRedirect)
}

// 모든 todoList → json으로 전송할 예정
func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

// 테스트용 임시 todos
func addTestTodos() {
	todoMap[1] = &Todo{1, "Study", true, time.Now()}
	todoMap[2] = &Todo{2, "Exercise", false, time.Now()}
	todoMap[3] = &Todo{3, "Play Game", false, time.Now()}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestTodos()
	rd = render.New()

	r := mux.NewRouter() // goriila/mux
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	return r
}
