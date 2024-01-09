package main

// 패키지명이 같을 때는, alias로 구분

import (
	htemplate "html/template"
	"text/template"
)

func main() {
	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)
}
