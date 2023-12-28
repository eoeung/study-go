// 테스트 할 때는 _test라고 붙이면, go가 알아서 인식
// 1) goconvey
//   - $ go install github.com/smartystreets/goconvey
//
// 2) testify/assert
//   - $ go get github.com/stretchr/testify/assert
package myapp

import (
	"io"
	"net/http"
	"net/http/httptest"
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
