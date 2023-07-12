package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8
	//byte 等同于int8，常用来处理ascii字符
	//
	//rune 等同于int32,常用来处理unicode或utf-8字符
	fmt.Println(utf8.RuneCountInString("你好"))
	fmt.Println(len([]rune("你好")))
	fmt.Println(len("你好"))
}
