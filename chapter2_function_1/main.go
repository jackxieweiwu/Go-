// chapter_function_1 project main.go
package main

import (
	"fmt"
)

//不定参数

//(1) 不定参数类型
// 下面这段代码的意思是，函数myfunc()接受不定数量的参数，这些参数的类型全部都是int
// 形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
//从内部实现机理来说，类型...type本质上是一个数组切片。也就是[]type，这也是为什么下面
//的参数可以用for循环来获得每个传入的参数

//
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

//(2) 不定参数的传递
func myfunc2(args ...int) {
	fmt.Println("only print from the second number")

	//传递片段，实际上任意的int slice都可以传递进去
	myfunc(args[1:]...)
}

//(3) 任意类型的不定参数
//如果需要传任意类型的可变参数，可以指定类型为interface{}。例如下面是Go语言标准库中的
//fmt.Printf()的函数原型
// func Printf(format string, args ...interface{})
//用interface{}传递任意类型数据是Go语言的惯例用法
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

//(4) 多返回值
//例如：func (file *FILE) Read(b []byte)(n int,err Error)
//从上面的方法原型可以看到，我们还可以给返回值命名，就像函数的输入参数一样。返回值被命名之后,
//它们的值在函数开始的时候被自动初始化为空。在函数中执行不带任何参数的return语句时，会返回
//对应的返回值变量的值
//Go语言并不需要强制命名返回值，但是命名后的返回值可以让代码更清晰，可读性更强，同事也可用于文档

//(5) 匿名函数和闭包
//5.1 匿名函数
// 在Go里面，函数可以像普通变量一样被传递或者使用，这与C语言的回调函数比较类似。不同的是Go语言里
// 支持随时在代码里定义匿名函数。
var AnoFunc = func(x, y int) int {
	return x + y
}

func main() {
	c, err := Add(2, 3)

	if err != nil {
		fmt.Println("parameter error")
		return
	}

	fmt.Println("result:", c)

	myfunc(1, 2, 3)
	fmt.Println("==================")
	myfunc(4, 5, 6)

	fmt.Println("======================")
	myfunc2(7, 8, 9, 10)

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	//调用匿名函数
	result := AnoFunc(20, 30)
	fmt.Println("result:", result)

	//调用匿名函数2
	f1 := func(a, b int) int {
		return a + b
	}
	result2 := f1(3, 4)
	fmt.Println("result2:", result2)

	//调用匿名函数3
	result3 := func(a, b int) int {
		return a + b
	}(5, 6)
	fmt.Println("result3:", result3)

	//闭包是可以包含自由变量的代码块，这些变量不在这个代码块内部或者任何全局上下文中定义，
	//而是在定义代码块的环境中定义。要执行的代码块为自由变量提供绑定的计算环境（作用域）
	var j int = 5
	a := func() func() {
		var i int = 10
		fmt.Println("POINT 1")
		return func() {
			fmt.Printf("i,j: %d %d\n", i, j)
		}
	}()

	fmt.Println("POINT 2")
	a()
	j *= 2
	a()

}
