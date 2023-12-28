package myapp

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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
	assert.Contains(string(data), "No User Id:89")

	fmt.Println(ts.URL + "/users/555555")
	resp2, err2 := http.Get(ts.URL + "/users/555555")
	assert.NoError(err2)
	assert.Equal(http.StatusOK, resp2.StatusCode)

	data2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(data2))
	assert.Contains(string(data2), "No User Id:555555")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ym", "last_name":"moon", "email":"eoeung113@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp2, err2 := http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err2)
	assert.Equal(http.StatusOK, resp2.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp2.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user, user2)
	assert.Equal(user.FirstName, user2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := io.ReadAll(resp.Body)
	log.Print(string(data))
	assert.Contains(string(data), "No User ID:1")

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"ym", "last_name":"moon", "email":"eoeung113@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ = io.ReadAll(resp.Body)
	fmt.Println(string(data))
	assert.Contains(string(data), "Deleted User ID:1")
	// assert.Contains(string(data), "No User ID:1")
}
