// JSON Transfer
// 1) QueryString
// 2) JSON

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type carHandler struct{}

type Car struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (c *carHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nc := new(Car)

	// print body
	// io.ReadAll을 같이 사용하면, json Decode에서 에러 발생
	// 이미 io.ReadAll(r io.Reader)로 r.Body를 읽어버리기 때문에
	// r.Body의 내용을 전체적으로 읽어와 버퍼에 저장한다.
	// → r.Body의 읽기 위치가 마지막으로 이동한다.
	body, _ := io.ReadAll(r.Body)
	fmt.Println(string(body))

	// print nc
	fmt.Println(nc)

	// 2) JSON
	// Request Body에서 받아온 값을 Decode → 문제 없는지 확인
	err := json.NewDecoder(r.Body).Decode(nc) // 포인터인 nc에, 인코딩된 JSON 값인 r.Body를 읽어 저장

	// print nc
	fmt.Println(nc)

	// 에러 처리
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error : ", err)
		return
	}

	// 다시 클라이언트에게 JSON으로 응답
	data, _ := json.Marshal(nc)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func name(w http.ResponseWriter, r *http.Request) {
	// 1) QueryString
	name := r.URL.Query().Get("name")
	fmt.Println(name)
	if name != "" {
		fmt.Fprintf(w, "name : %s", name)
		return
	}

	fmt.Fprintf(w, "name값이 없습니다.")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", name)         // 1) QueryString
	mux.Handle("/car", &carHandler{}) // 2) JSON

	http.ListenAndServe(":3000", mux)
}
