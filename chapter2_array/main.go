// chapter2_array project main.go
package main

import (
	"fmt"
)

func modify(array [5]int) {
	array[0] = 100

	fmt.Println("In modify,array values:", array)
}

func main() {

	var a [32]byte

	fmt.Println("length of a:", len(a))

	//(1) 数组访问
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("=========================")

	//Go语言还提供了一个关键字range，用于便捷的遍历容器中的元素。当然，数组也是
	//range的支持范围。上面的遍历过程可以简化为如下的写法：
	for i, v := range a {
		fmt.Printf("%d ", v)
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}

	//(2) 值类型
	//需要特别注意的是，在Go语言中数组是一个值类型（valuetype）。所有的值类型变量在赋值和作为参数传递时，
	//都将产生一次赋值动作

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	modify(b)
	fmt.Println("In main,array values:", b)

	//(3) 遍历
	var c [5]int = [5]int{100, 101, 102, 103, 104}
	for i, v := range c {
		fmt.Println("index:", i, " value:", v)
	}

}
