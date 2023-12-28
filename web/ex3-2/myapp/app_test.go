// 테스트 할 때는 _test라고 붙이면, go가 알아서 인식
// 1) goconvey
//   - $ go install github.com/smartystreets/goconvey
//
// 2) testify/assert
//   - $ go get github.com/stretchr/testify/assert
package myapp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	// 실제 동작은 잘 되는데, 에러나는 이유
	// 계속 GOPATH, GOROOT에서 해당 패키지를 찾으려 한다.
)

func TestIndexPathHandler(t *testing.T) {
	// assert.New(t) : 지정된 t에 대한 새로운 Assertions 객체 생성
	assert := assert.New(t)

	res := httptest.NewRecorder() // 테스트용 response
	// *httptest.ResponseRecorder를 반환
	req := httptest.NewRequest("GET", "/", nil) // 테스트용 request // method, url, request body
	// *http.Request 반환

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// indexHandler(res, req) // 같은 패키지라서 소문자라도 인식

	assert.Equal(http.StatusOK, res.Code) // 200 status code와 res.Code가 일치하는지 비교
	// return &ResponseRecorder{
	// 	HeaderMap: make(http.Header),
	// 	Body:      new(bytes.Buffer),
	// 	Code:      200,
	// }
	// 기본적으로 200 code로 초기화됨

	// assert.Equal과 같은 코드
	// if res.Code == http.StatusOK {
	// 	t.Fatal("Success!! ", res.Code)
	// }

	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// barHandler(res, req) // barHandler 자체에는 API 호출 주소가 없음

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=영희", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	// barHandler(res, req) // barHandler 자체에는 API 호출 주소가 없음

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello 영희", string(data))
	// assert.Equal("Hello World", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code) // fooHandler에서 http.StatusBadRequest로 설정
	// json.NewDecoder(r.Body).Decode(user) 에서 r.Body가 nil이기 때문에 에러 발생
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"FirstName":"ym", "LastName":"moon", "Email":"eoeung113@gmail.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	// 1) NewDecoder로 res.Body에서 JSON 디코딩 수행하는 Decoder 생성
	// 2) Decode는 생성한 디코더를 이용해서 JSON 데이터를 읽고, user에 매핑
	// 3) res.Body에서 읽어온 JSON 데이터가 구조체 필드에 매핑되어 user 인스턴스에 저장됨
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("ym", user.FirstName)
	assert.Equal("moon", user.LastName)
}
