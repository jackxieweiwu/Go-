// chapter5_json project main.go
package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	gobook := &Book{
		Title:       "go语言编程",
		Authors:     []string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan"},
		Publisher:   "ituring.com.cn",
		IsPublished: true,
		Price:       9.99,
	}

	b, err := json.Marshal(gobook)
	if err != nil {
		fmt.Println("Marshal error")
		return
	}

	var str string = string(b)
	fmt.Println("str:", str)

	var unmarshalBook *Book
	unmarshalBook = new(Book)
	err = json.Unmarshal(b, unmarshalBook)
	if err != nil {
		fmt.Println("Unmarshal error")
		return
	}
	fmt.Println("unmarshalBook:", unmarshalBook)
}
