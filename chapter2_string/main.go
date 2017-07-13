// chapter2_string project main.go
package main

import (
	"fmt"
)

func main() {

	var str string
	str = "Hello,World" //字符串赋值
	ch := str[0]        //取字符串的第一个字符
	fmt.Printf("The length of \"%s\" is %d\n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	//注意，与C语言中的字符串类型相似，字符串的内容不能在初始化后被修改（属于字符串常量）

	//字符串操作
	//(1) 字符串连接
	var x, y string
	x, y = "hello", "123"
	fmt.Println("x+y=", x+y)

	//(2) 字符串长度
	fmt.Println("len(x):", len(x), "len(y)", len(y))

	//(3) 取字符
	fmt.Printf("Hello[1] is %c\n", "Hello"[1])

	//(4) 字符串遍历

	//每个中文字符在UTF-8中占用3个字节，而不是一个字节
	str = "你好，世界"
	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	fmt.Println("======================================")
	str = "Hello,世界"
	//以Unicode字符遍历.以Unicode字符方式遍历时，每个字符类型是rune
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}

	//在Go语言中支持两个字符类型，一个是byte(实际上是uint8的别名），代表UTF-8字符串的单个字节的值，
	//另一个是rune，代表的是单个Unicode字符
}
