// chapter_slice project main.go
package main

import (
	"fmt"
)

func modify(array []int) {
	array[0] = 100

	fmt.Println("In modify,array values:", array)
}

//数组切片的数据结构可以抽象为以下3个变量
//(1) 一个指向原生数组的指针
//(2) 数组切片中的元素个数
//(3) 数组切片已分配的存储空间
//从底层实现的角度来看，数组切片实际上仍然使用数组来管理，因此它们之间的关系让C++程序员
//很容易联想到STL中的std::vector和数组的关系。基于数组，数组切片添加了一系列管理功能
//可以随时动态扩充空间，并且可以被随意传递而不会导致所管理的元素被重复复制

func main() {

	//1: 创建数组切片
	//创建数组切片的方法主要有两种：基于数组和直接创建

	//(1) 基于数组
	//数组切片可以基于一个已存在的数组创建。数组切片可以只使用数组的一部分元素或整个数组来创建

	//定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//基于数组创建一个数组切片(表示的范围是[0,5)个元素，即myArray[0],myArray[1],myArray[2],myArray[3],myArray[4]
	var mySlice []int = myArray[3:5]

	fmt.Println("Elements of myArray:")
	for _, v := range myArray {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice:")

	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println()

	//对所有数组元素切片
	mySlice = myArray[:]
	modify(mySlice) //这样就可以修改数组元素的值了
	fmt.Println("myArray:", myArray)

	//（2） 直接创建切片
	// go语言提供的内置函数make()可以用于灵活地创建数组切片

	//创建一个初始元素个数为5的数组切片，元素初始值为0
	mySlice1 := make([]int, 5)
	for _, v := range mySlice1 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice2 := make([]int, 5, 10)
	for _, v := range mySlice2 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//直接创建并初始化包含5个元素的数组切片
	mySlice3 := []int{1, 2, 3, 4, 5}
	for _, v := range mySlice3 {
		fmt.Print(v, " ")
	}
	fmt.Println()

	//2: 数组元素的遍历

	//3:动态增减元素
	mySlice4 := make([]int, 5, 10)

	fmt.Println("len(mySlice4):", len(mySlice4))
	fmt.Println("cap(mySlice4):", cap(mySlice4))

	//往mySlice4中已包含的5个元素后面继续新增元素
	mySlice4 = append(mySlice4, 6, 7)
	fmt.Println("mySlice4:", mySlice4)

	mySlice5 := []int{8, 9, 10, 11, 12, 13, 14}
	mySlice4 = append(mySlice4, mySlice5...) //这里...的意思是把mySlice5包含的所有元素打散后传入
	fmt.Println("mySlice4:", mySlice4)

	//4: 基于数组切片创建数组切片
	// 下面创建切片时，选择的oldSlice元素范围甚至可以超过所包含的元素个数。只要不超过oldSlice的cap()
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[3:]
	fmt.Println("newSlice:", newSlice)

	//5: 内容复制
	//如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	slice3 := []int{1, 2, 3, 4, 5}
	copy(slice1, slice2)
	copy(slice2, slice3)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	//6:删除元素
	fmt.Println("delete elements")
	deslice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("delete 3")
	for i, v := range deslice1 {
		if v == 3 {
			if len(deslice1) == 1 {
				deslice1 = make([]int, 0)
			} else if i == len(deslice1)-1 {
				deslice1 = deslice1[:i]
			} else if i == 0 {
				deslice1 = deslice1[1:]
			} else {
				deslice1 = append(deslice1[:i], deslice1[i+1:]...)
			}
			break
		}
	}

	for _, v := range deslice1 {
		fmt.Printf("%d ", v)
	}
}
