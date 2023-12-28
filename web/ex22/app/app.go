package app

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"web/ex22/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render = render.New()
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

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
	// Todo리스트를 가져올 수 있다는 것은, 이미 로그인이 되었다는 말과 동일
	// → getSessionID() 함수로 sessionId값을 가져올 수 있음
	sessionId := getSessionID(r)
	// todoMap에 저장되어 있는 모든 *Todo(포인터 변수)를 list라는 슬라이스에 저장 후 반환
	list := a.db.GetTodos(sessionId)
	rd.JSON(w, http.StatusOK, list)
}

// todo 추가
func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	// 마찬가지로 Todo를 Insert할 수 있다는 것은, 이미 로그인이 되었다는 말과 동일
	// → getSessionID() 함수로 sessionId값을 가져올 수 있음
	sessionId := getSessionID(r)
	name := r.FormValue("name")
	todo := a.db.AddTodo(sessionId, name)
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

// Session에 저장된 ID를 가져온다.
// function 포인터를 갖는 변수를 생성
// → 기존에 func getSessionID로 작성했지만, 테스트 코드 실행하자고 함부로 바꾸면 위험해짐
var getSessionID = func(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}
	// Set some session values.
	val := session.Values["ID"]
	if val == nil {
		return ""
	}
	// 타입 어순 변환 (type assertion)
	return val.(string) // string타입이라고 가정 → val값을 string으로 형변환
}

// 로그인 체크
func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// request URL이 /signin.html이면, next()를 해야함
	// → /signin.html로 요청했는데, 미들웨어에서는 다시 로그인 체크를 시도 → /signin.html 요청
	// → 무한 반복 가능성 존재
	if strings.Contains(r.URL.Path, "/signin") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r) // 현재 미들웨어에서 추가적인 처리가 필요하지 않고, 다음 미들웨어 핸들러로 제어를 전달
		return
	}

	sessionID := getSessionID(r)

	// 유저가 로그인 되어있으면 다음 핸들러로 이동
	if sessionID != "" {
		next(w, r)
		return
	}

	// 유저가 로그인 되어있지 않으면, redirect 시켜야한다.
	// → signin.html로 이동
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter() // goriila/mux
	// 미들웨어 등록
	// negroni.Classic()은 default값으로 이미 미들웨어를 등록했기 때문에, custom으로 설정
	// 데코레이터 패턴 적용 → 첫번째 인자부터 순서대로 적용한다.
	// 4가지 미들웨어 중에서, 첫번째 부터 시작하는데,
	// 미들웨어를 만족하지 못하면 거기서 다음 미들웨어로 가지 않고 종료하고, 다음 코드로 이동
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	// AppHandler : 핸들러/DB를 모두 사용할 수 있도록 설정
	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	// 구글 로그인
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)

	return a
}

// ex17에서 진행했던 todoMap을 이용해서 데이터를 저장하는 방식을 분리하기 위해 Refactoring 진행
// → 데이터를 메모리에 가지고 있는 것이 아닌, 파일 DB로 분리/저장
// → 의존성을 깨뜨리기위한 작업

// SQLite : RDBMS, Query, 무료
