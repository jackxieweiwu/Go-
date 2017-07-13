// chapter2_controlflow project main.go
package main

import (
	"fmt"
)

//条件语句,对应的关键字为if,else,和else if
//选择语句，对应的关键字为switch,case和select
//循环语句，对应的关键字for,range
//跳转语句，对应的关键字为goto

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
	//fmt.Println("run here in example")			//此处不会跑到这里，打开这行语句将导致失败
}

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}

func main() {

	//1: 条件语句应该注意以下几点
	//(1)条件语句不需要使用括号将条件包含起来()
	//(2) 无论语句体内有几条语句，花括号{}都必须存在
	//(3) 左花括号{必须与if或者else处于同一行
	//(4) 在if之后，条件语句之前，可以添加变量初始化语句，使用；间隔
	//(5) 在有返回值的函数中，不允许将“最终的” return语句包含在if...else...结构中
	fmt.Println()

	value := example(30)
	fmt.Printf("value:%d\n", value)

	//2:switch语句 (fallthrough关键字，不用break语句）
	value2 := 7

	switch value2 {
	case 0:
		fmt.Println("value is 0")
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
		//	case 3:								//Here is error
		//	case 4:
		//		fmt.Println("value is 3 or 4")
	case 3, 4:
		fmt.Println("value is 3 or 4")
	case 5:
		fallthrough
	case 6:
		fmt.Println("value is 5 or 6")
	case 7, 8, 9: //单个case中可以出现多个结果选项
		fmt.Println("value is 7 or 8 or 9")
	default:
		fmt.Println("error")
	}

	//3: 循环语句
	//与多数语言不同的是，Go语言中的循环语句只支持for关键字，而不支持while和do-while
	//(1) 左花括号{必须与for处于同一行
	//(2) Go语言中的for循环与C语言一样，都允许在循环条件中定义和初始化变量，唯一的区别是，
	//    Go 语言不支持以逗号为间隔的多个赋值语句，必须使用平行赋值的方式来初始化多个变量
	//(3) Go语言for循环同样支持continue和break来控制循环，但是它提供了一个更高级的break
	//    可以选择中断哪一个循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	//无限循环的一种简单写法
	sum = 0
	for {
		sum++
		if sum >= 100 {
			break
		}
	}
	fmt.Println("sum:", sum)

	//在条件表达式中也支持多重赋值
	a := []int{1, 2, 3, 4, 5, 6}

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	fmt.Println("a:", a)

	myfunc()
}
