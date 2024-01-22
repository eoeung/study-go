package AwesomeDB

import "fmt"

type OurDB struct {
	Name string
}

func (db *OurDB) GetData() string {
	return "GetData()"
}

func (db *OurDB) WriteData(data string) {
	fmt.Println("WriteData()")
}

func (db *OurDB) Close() error {
	return nil
}
