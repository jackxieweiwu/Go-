// variable_test project main.go
package main

import (
	"fmt"
)

//测试返回值
func my_print(a int) int {
	fmt.Printf("a = %d\n", a)

	return a
}

//go语言中提供了C/C++程序员期盼多年的多重赋值功能
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

//匿名变量
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi maruko"
}

//常量的定义
const Pi float64 = 3.1415926
const zero = 0.0 //无类型浮点型常量
const (
	size int64 = 1024
	eof        = -1 //无类型整形常量
)
const cu, cv float32 = 0, 3 //cu=0.0,cv=3.0,常量的多重赋值

const cd, ce = 20, 30.2 //无类型常量的另一种定义方法

const ci, cj, ck = 3, 4, "foo" //ci=3,cj=4,ck="foo"，无类型整形和字符串常量

const mask = 1 << 3

//const Home = os.GetEnv("HOME")  //常量的赋值是一个编译期的行为，所以右值不能出现任何需要运行期才能得出结果的表达式

func main() {
	//variable declare
	var v1 int
	var v2 string
	//var v3 [10]int //array
	//var v4 []int   //slice
	//var v5 struct {
	//	f int
	//}
	//var v6 *int            //pointer
	//var v7 map[string]int  //map,the 'key' is string type and value is int type
	var v8 func(a int) int //function pointer type
	var result int

	v1 = 20
	fmt.Println("v1 =", v1)

	v2 = "Hello,world"
	fmt.Println("v2:", v2)

	v8 = my_print
	result = v8(40)
	fmt.Printf("result:%d\n", result)

	// variable initialize
	var v9 int = 10 //正确使用方式1
	var v10 = 100   //正确使用方式2，编译器可以自动推导出v2的类型
	v11 := 1000     //正确使用方式3，编译器可以自动推导出v3的类型（用于明确表达同事进行变量声明和初始化的工作）
	fmt.Printf("v9:%d v10:%d v11:%d\n", v9, v10, v11)

	var (
		a int
		b int
	)
	a = 10
	b = 20
	fmt.Printf("before swap: a %d b %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap: a %d b %d\n", a, b)

	var (
		a1 int = 100
		b1 int = 200
	)
	fmt.Printf("a1:%d b1:%d\n", a1, b1)

	var c, d float32
	c = 1.01
	d = 2.01
	fmt.Printf("c:%f d:%f\n", c, d)

	//The following is Error,can't use this style(不能在声明变量的时候又进行多重赋值)
	//var ef, ff float32 = 3.1, 3, 2
	//fmt.Printf("e:%f f:%f\n", ef, ff)

	e, f := 3.1, 3.2
	fmt.Printf("e:%f f:%f\n", e, f)

	//call getName function
	_, _, nickName := GetName()
	fmt.Printf("nickName:%s\n", nickName)

	//常量的打印
	fmt.Printf("cd:%d ce:%f\n", cd, ce)
	fmt.Printf("ci:%d cj:%d ck:%s\n", ci, cj, ck)
	fmt.Printf("mask:%d\n", mask)
}
