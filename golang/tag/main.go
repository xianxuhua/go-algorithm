package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	// 解析tag
	u := &User{}
	name, ok := reflect.TypeOf(u).Elem().FieldByName("Name")
	if !ok {
		panic("")
	}
	fmt.Println(string(name.Tag))
}
