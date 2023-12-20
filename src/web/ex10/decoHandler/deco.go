package decoHandler

import (
	"fmt"
	"net/http"
	"reflect"
)

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

// http.Handler 인터페이스를 구현해야함
func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(reflect.TypeOf(self.h))
	self.fn(w, r, self.h)
}

/*
	구현하지 않으면 아래 에러 발생
	cannot use &DecoHandler{…} (value of type *DecoHandler) as http.Handler value in return statement:
	*DecoHandler does not implement http.Handler (missing method ServeHTTP)
*/

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h, // wrapping : 기존 핸들러(http.Handler)를 감싸고 있음
	}
}
