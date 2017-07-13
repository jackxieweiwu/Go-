// chapter2_map project main.go
package main

import (
	"fmt"
)

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {

	var personDB map[string]PersonInfo

	personDB = make(map[string]PersonInfo)

	//往这个personDB里面插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101..."}

	fmt.Println("personDB:", personDB)

	//从这个map中查找键为“1234”的信息
	person, ok := personDB["1234"]

	//ok是一个返回的bool型，返回TRUE表示找到了对应的数据
	if ok {
		fmt.Println("Found Person", person.Name, "with ID 1234")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}

	//使用map的步骤
	//(1) 变量声明
	var myMap map[string]PersonInfo
	var myMap2 map[string]PersonInfo
	var myMap3 map[string]PersonInfo

	//(2) 创建
	//方式1
	myMap = make(map[string]PersonInfo)
	myMap["3333"] = PersonInfo{"3333", "test", "NULL"}

	//方式2(指定存储能力）
	myMap2 = make(map[string]PersonInfo, 100)
	myMap2["3333"] = PersonInfo{"444", "test2", "NULL"}

	//方式3(创建并初始化)
	//这里需要注意的两个地方：1-- key与value之间以冒号分割 2--末尾有一个逗号
	myMap3 = map[string]PersonInfo{
		"4567": PersonInfo{"4567", "test3", "NULL"},
		"8910": PersonInfo{"8910", "test4", "NULL"},
	}

	fmt.Println("myMap:", myMap)
	fmt.Println("myMap2:", myMap2)
	fmt.Println("myMap3:", myMap3)

	//(3) 元素赋值
	myMap["39390"] = PersonInfo{"39390", "test_1", "NULL"}

	//(4) 元素删除
	//删除时，如果传入的的键值不存在，那么这个调用将什么都不发生。但是如果传入的map变量为nil，该调用
	//将导致程序抛出异常
	fmt.Println("myMap:", myMap)
	delete(myMap, "3333")
	fmt.Println("myMap:", myMap)

	var myMap4 map[string]PersonInfo
	if myMap4 == nil {
		fmt.Println("currently,myMap4 is nil")
	} else {
		fmt.Println("currently,myMap4 is not nil")
	}

	//(5) 元素查找
	// value, ok := myMap["1234"]
}
