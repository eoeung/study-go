package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

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
	rd.JSON(w, http.StatusOK, user)

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
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)

		// unrolled/render 사용
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()

	// unrolled/render 사용
	rd.JSON(w, http.StatusOK, user)

	// w.Header().Add("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError)
	// 	// fmt.Fprint(w, err)

	// 	rd.Text(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "ymmoon")

	user := User{Name: "ymmoon", Email: "eoeung113@gmail.com"}

	// unrolled/render
	// 기존 템플릿 사용 + rd.HTML 기본값 사용 (디렉토리 : templates / 확장자 : tmpl)
	rd.HTML(w, http.StatusOK, "hello", "ymmoon") // 기본적으로 templates라는 디렉토리안에 있는 tmpl 확장자 파일을 읽어들임

	// rd.HTML 기본값이 아닌 customize 값 사용 (디렉토리 : template / 확장자 : html)
	rd2.HTML(w, http.StatusOK, "body", user) // tmpl이 아닌 다른 확장자를 읽어들이기 위해서는 옵션을 지정해줘야 함
}

func main() {
	rd = render.New()

	options := render.Options{
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello2",
	}
	rd2 = render.New(options)
	// gorilla/pat 사용
	mux := pat.New()
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// mux.Handle("/", http.FileServer(http.Dir("public")))
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
