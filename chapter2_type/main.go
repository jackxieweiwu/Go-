// chapter2_type project main.go
package main

import (
	"fmt"
	"math"
)

//p为用户定义的比较精度，比如0.0001
func IsEqual(f1, f2, p float64) bool {
	return math.Abs(f1-f2) < p
}
func main() {

	//(1) bool 类型
	var v1 bool = true
	fmt.Println("Result:", v1)
	v2 := (1 == 2) //v2也会被推导出bool类型
	fmt.Println("v2:", v2)

	//bool类型不能接受其他类型的赋值，不支持自动或强制类型转换。如下的示例是
	//一些错误的用法，会导致编译错误
	//var b = bool
	// b = 1		//编译错误
	// b = bool(1)  //编译错误

	//(2) 整形
	//int8 uint8 int16 uint16
	//uintptr 同指针，在32位平台下为4字节，64位平台下位8字节

	var ia uintptr = 2048
	fmt.Printf("ia:%d\n", ia)

	//2.1 需要注意的是，int和int32在Go语言里被认为是两种不同的类型，编译器也不会帮你
	//自动做类型转换
	//var value2 int32
	//value1 := 64		//value1被自动推导为Int类型
	//value2 = value1   //编译错误
	//value2 = int32(value1)  //编译通过

	//2.2比较运算
	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal")
	} else {
		fmt.Println("i and j are not equal")
	}
	//两个不同类型的整型数不能直接比较，比如int8类型的数和int类型的数不能直接进行比较。
	//但是各种类型的整型变量都可以直接与字面量(literal)进行比较
	var i2 int32
	var j2 int64

	i2, j2 = 2, 3

	//		if i2 == j2 {
	//			fmt.Println("i2 and j2 are equal")
	//		} else {
	//			fmt.Println("i2 and j2 are not equal")
	//		}

	if i2 == 2 || j2 == 3 {
		fmt.Println("i2 and j2 are equal")
	} else {
		fmt.Println("i2 and j2 are not equal")
	}

	//go语言的大多数位运算都与C语言比较类似，除了取反在C语言中是~x，而在Go语言中是^x。例如
	var x2 = ^2
	fmt.Println("x2:", x2)

	//2.3 浮点类型
	//go语言定义了两个类型float32和float64,其中float32等价于C语言的float类型，float64等价于C语言的double类型
	var fvalue1 float32
	fvalue1 = 12

	//如果不加小数点，fvalue2会被推导位整型而不是浮点型（需要注意的是，其类型将被自动推导为float64,
	//而不管赋给它的数字是否是用32为长度表示的
	fvalue2 := 12.0

	//fvalue1 = fvalue2
	fvalue1 = float32(fvalue2)
	fmt.Println("fvalue1:", fvalue1)
	fmt.Println("fvalue2:", fvalue2)

	//浮点数的比较
	fmt.Println("is 2.0 and 2.0000001 equal?", IsEqual(2.0, 2.0000001, 0.0000001))

	//2.4 复数类型
	var value_complex complex64 //由两个float32构成的复数类型

	value_complex = 3.2 + 12i
	value_complex2 := 3.2 + 13i        //value_complex2是complex128类型
	value_complex3 := complex(3.2, 13) //value3结果同value2

	fmt.Println("value_complex:", value_complex)
	fmt.Println("value_complex2:", value_complex2)
	fmt.Println("value_complex3:", value_complex3)
	fmt.Printf("value_complex3 real:%f imag:%f\n", real(value_complex3), imag(value_complex3))
}
