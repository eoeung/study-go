package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := io.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	// resp, err := http.Get(ts.URL)
	resp, err := http.Get(ts.URL + "/ttt")
	fmt.Printf("main_test :::::::::: %s\n", ts.URL+"/ttt")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(buf)    // binary buffer를 한 줄로 읽기 위해 bufio 패키지 호출
	line, _, err := r.ReadLine() // 한 줄씩 읽는 bufio.Reader의 메서드
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER2] Started")

	// 한 줄 더 읽는다.
	line, _, err = r.ReadLine() // 한 줄씩 읽는 bufio.Reader의 메서드
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")
}
