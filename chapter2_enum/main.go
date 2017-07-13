// chapter2_enum project main.go
package main

import (
	"fmt"
)

//枚举
//枚举指一系列相关的常量，比如下面关于一个星期中每天的定义
//Go 语言并不支持众多其他语言明确支持的enum关键字

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays //这个常量没有导出
)

//说明：同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见.
//上面例子中numberOfDays为包内私有，其他符号则可被其他包访问

func main() {
	fmt.Printf("numberofDays:%d\n", numberOfDays)
}
