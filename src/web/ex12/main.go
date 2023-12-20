package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

// [Render 구조체]
// http response에 json, xml, binary data, html 템플릿을 쉽게 작성하는 함수를 제공
var rd *render.Render
var rd2 *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "ymmoon", Email: "eoeung113@gmail.com"}

	// unrolled/render 사용
	rd.JSON(w, http.StatusOK, user) // json.Marshal()과 동일한 기능

	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	// 클라이언트에서 보내준 정보를 바탕으로 유저 생성
	// → json으로 되어있는 정보를 디코딩
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// 기존 방법
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)

		// unrolled/render 사용
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()

	// unrolled/render 사용
	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 기존 방법
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError) // = 같은 구문 rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	// fmt.Fprint(w, err)
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "ymmoon")

	user := User{Name: "ymmoon", Email: "eoeung113@gmail.com"}

	// unrolled/render
	// 기존 템플릿 사용 + rd.HTML 기본값 사용 (디렉토리 : templates / 확장자 : tmpl)
	// "ymmoon" 문자열은 hello.tmpl이라는 템플릿에 넣을 값
	rd.HTML(w, http.StatusOK, "hello", "rd.HTML() ::: {{.}}의미 : 여기에 넣은 값이 그대로 출력됨") // 기본적으로 templates라는 디렉토리안에 있는 tmpl 확장자 파일을 읽어들임

	// rd.HTML 기본값이 아닌 customize 값 사용 (디렉토리 : template / 확장자 : html)
	// "body.html"을 메인으로 잡고 설정
	// "body.html"은 {{.Name}}, {{.Email}}이라는 Name, Email필드를 찾고 있으므로, string으로 보내주면 안됨
	rd2.HTML(w, http.StatusOK, "body", user) // tmpl이 아닌 다른 확장자를 읽어들이기 위해서는 옵션을 지정해줘야 함
}

func main() {
	rd = render.New() // 새로운 Render 인스턴스 생성

	options := render.Options{
		Directory:  "template",                 // default : "templates"
		Extensions: []string{".html", ".tmpl"}, // default : [".tmpl"]
		Layout:     "layout_hello",             // 레이아웃 설정 시, hello2 안에 있는 {{ yield }}에 값이 들어감
		// → 이 Layout이 템플릿이 되고, 추후에 설정한 값이나 템플릿을 이 Layout안에 넣는다.
	}
	rd2 = render.New(options) // 새로운 Render 인스턴스 생성
	// gorilla/pat 사용
	mux := pat.New() // *Router 반환
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// mux.Handle("/", http.FileServer(http.Dir("public")))
	n := negroni.Classic()

	/*
		Classic()
		func Classic() *Negroni {
			return New(NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
		}

		// NewRecovery() *Recovery
		// NewLogger() *Logger
		// NewStatic(directory http.FileSystem) *Static
		// http.Dir은 http.FileSystem 인터페이스를 구현하고 있음
	*/
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
