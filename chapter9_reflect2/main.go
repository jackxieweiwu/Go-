// chapter9_reflect2 project main.go
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) //注意得到的是X的地址
	fmt.Println("type of p:", p.Type())

	fmt.Println("settability of p:", p.CanSet())

	v := p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.2)
	fmt.Println(v.Interface())
	fmt.Println(x)

	//-----------------
	t := T{20, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
