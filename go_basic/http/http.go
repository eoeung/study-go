package main

import (
	"net/http"
)

func main() {
	http.Handle("/", new(testHandler))

	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Request Path is " + req.URL.Path
	w.Write([]byte(str))
}

/*
	＊"/" 요청을 보낼 때, ServeHTTP가 호출되는 이유

		1) http.ListenAndServe(":5000", nil)
		2) return server.ListenAndServe()
		3) return srv.Serve(ln)
		4) go c.serve(connCtx)
		5) serverHandler{c.server}.ServeHTTP(w, w.req)
*/
