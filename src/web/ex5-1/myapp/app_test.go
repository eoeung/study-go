package myapp

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	fmt.Println(ts.URL) // http://127.0.0.1:53040 → 생성한 testServer의 주소 같음
	assert.NoError(err) // 에러를 반환하지 않을 때 true / 에러를 반환하면 Fail() 실행
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close() // 모든 작업이 끝나면 서버 셧다운

	// Get 요청
	fmt.Println(ts.URL + "/users")
	resp, err := http.Get(ts.URL + "/users")
	// mux에 '/users'를 등록하지 않았음에도 test에서 Pass됨
	// → /users가 없는 경우, 상위 핸들러가 호출됨
	// => /가 호출됨 // mux.HandleFunc("/", indexHandler)이 호출됨
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := io.ReadAll(resp.Body)
	// assert.Equal("Hello World", string(data)) // 실제로 /와 /users가 구별되는지 확인
	fmt.Println(string(data))
	assert.Contains(string(data), "Get UserInfo")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	fmt.Println(ts.URL + "/users/89")
	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := io.ReadAll(resp.Body)
	fmt.Println(string(data))
	assert.Contains(string(data), "User Id:89")

	fmt.Println(ts.URL + "/users/555555")
	resp2, err2 := http.Get(ts.URL + "/users/555555")
	assert.NoError(err2)
	assert.Equal(http.StatusOK, resp2.StatusCode)

	data2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(data2))
	assert.Contains(string(data2), "User Id:555555")
}
