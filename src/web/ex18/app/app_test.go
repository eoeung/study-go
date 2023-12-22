package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"web/ex18/model"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(MakeHandler()) // 핸들러를 넣어준다.
	defer ts.Close()

	// {"name": {"Test todo"}}이라고 해도 되나, 가독성을 위해 []string 추가
	// addTodoHandler()에서 r.FormValue("name")로 값을 받기 때문에, PostForm 사용
	// 새로운 유저1 등록
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": []string{"Test todo1"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// JSON으로 응답한 값 읽어오기
	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo1")
	id1 := todo.ID

	// 새로운 유저2 등록
	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": []string{"Test todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo2")
	id2 := todo.ID

	// 모든 todo 가져오기
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	var todos = []*model.Todo{}
	// *Todo 포인터 변수를 담는 슬라이스에, resp.Body값 매핑
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2) // 현재까지 두 명의 유저를 등록했음

	// 배열과 슬라이스의 range는 idx, val 반환
	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("Test todo1", t.Name)
		} else if t.ID == id2 {
			assert.Equal("Test todo2", t.Name)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	// 완료 여부 변경
	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?complete=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// 모든 todo 가져오기
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	// *Todo 포인터 변수를 담는 슬라이스에, resp.Body값 매핑
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2) // 현재까지 두 명의 유저를 등록했음

	// 배열과 슬라이스의 range는 idx, val 반환
	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	// todo 삭제
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// DELETE 됐는지 확인
	// 모든 todo 가져오기
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	// *Todo 포인터 변수를 담는 슬라이스에, resp.Body값 매핑
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 1) // 2명 등록 - 1명 삭제
	// 배열과 슬라이스의 range는 idx, val 반환
	for _, t := range todos {
		assert.Equal(t.ID, id2) // id1을 삭제했기 때문에, id2만 남아야한다.
	}
}
