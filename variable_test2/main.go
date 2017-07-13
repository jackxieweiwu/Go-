// variable_test2 project main.go
package main

import (
	"fmt"
)

//预定义常量
//Go语言中预定了这些常量:true,false和iota
//iota比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被重置为0，然后
//在下一个const出现之前，每出现一次iota，其所代表的数字会自动增1
const ( //iota被重置为0
	c0 = iota //c0 == 0
	c1 = iota //c1 == 1
	c2 = iota //c2 == 2
)

const (
	a = 1 << iota //a == 1(iota在每个const开头被重置为0）
	b = 1 << iota //b == 2
	c = 1 << iota //c == 4
)

const (
	u         = iota * 42 //u == 0
	v float64 = iota * 42 //v == 42.0
	w         = iota * 42 //w = 84

)

//如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。因此，上面的前两个const语句可以简写为：
const (
	c00 = iota //c00 == 0
	c01        //c01 == 1
	c02        //c02 == 2
)

const (
	a_1 = 1 << iota //a_1 = 1
	b_1             //b_1 = 2
	c_1             //c_2 = 4
)

const (
	a_2 = 2 //a_2 == 2
	b_2     //b_2 == 2
	c_2     //c_2 == 2
)

const x = iota //x == 0
const y = iota //y == 0

//const just_for_test int			//Error,must set the value
const just_for_test int = 20

var just_for_test2 int //this is OK

func main() {

	fmt.Printf("c2:%d\n", c2)
	fmt.Printf("c:%d\n", c)
	fmt.Printf("w:%d\n", w)
	fmt.Printf("c02:%d\n", c02)
	fmt.Printf("c_1:%d\n", c_1)
	fmt.Printf("b_2:%d c_2:%d\n", b_2, c_2)

	fmt.Printf("just_for_test:%d\n", just_for_test)
	fmt.Printf("just_for_test2:%d\n", just_for_test2)
}
