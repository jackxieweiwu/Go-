// chapter3_interface project main.go
package main

import (
	"fmt"
)

//在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。
//1:接口赋值
//接口赋值在Go语言中可以分为如下两种情况：
//1) 将对象实例赋值给接口
//2) 将一个接口赋值给另外一个接口

// 将对象实例赋值给一个接口
type Integer int

func (a Integer) Less(b Integer) bool { //
	return a < b
}

//可以生成一个下面的方法
//func (a *Integer) Less(b Integer) bool {
//	return (*a).Less(b)
//}
func (a *Integer) Add(b Integer) {
	*a += b
}

//不能生成下面的方法
//func (a Integer) Add(b Integer){
//	(&a).Add(b)
//}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

//将一个接口赋值给另外一个接口
// 在Go语言中只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值。

//2:接口查询
//接口查询语法（例如：）
//type IStream interface{
//	Write(buf []byte)(n int,err error)
//	Read(buf []byte)(n int,err error)
//}
//type Writer interface{
//	Write(buf []byte)(n int,err error)
//}
//var file1 Writer
//if file2,ok := file1.(IStream);ok{

//}

//3:类型查询

type Stringer interface {
	String() string
}

type MyString string

func (str MyString) String() string {
	return string(str)
}

func MyPrintln(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int: //现在v的类型是int
		case string: //现在v的类型是string
		default:
			if v, ok := arg.(Stringer); ok {
				val := v.String()
				fmt.Println("val:", val)
			}
		}
	}
}

//4:接口组合
//可以认为接口组合是类型匿名组合的一个特定场景，只不过接口只包含方法，而不包含任何成员变量

//5: Any类型
//由于Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向任何对象的Any类型，如下：
//var v1 interface{} = 1       //将int类型赋值给interface{}
//var v2 interface{} = "abc"   //将string类型赋值给interface{}
//var v3 interface{} = &v2     //将*interface{}类型赋值给interface{}
//var v4 interface{} = struct{X int}{1}
//var v5 interface{} = &struct{X int}{1}

func main() {
	var a Integer = 1
	var b LessAdder = &a

	fmt.Println(b.Less(20))
	b.Add(2)
	fmt.Println("1 + 2 =", a)

	var str MyString = "你好，世界"
	MyPrintln(10, 20, 1.2, "Hello,world", str)
}
