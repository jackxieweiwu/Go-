// chapter3_type project main.go
package main

import (
	"fmt"
)

//1:为类型定义方法
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

//2: 值语义和引用语义
//Go语言中的大多数类型都是基于值语义，包括：
//基本类型，如byte,int,bool,float32,float64和string等
//复合类型，如数组(array),结构体(struct)和指针(pointer)等

//Go语言中有4个类型比较特殊，看起来像引用类型，如下图所示：
//1) 数组切片：指向数组(array)的一个区间
//2)map：极其常见的数据结构，提供键值查询能力
//3)channel:执行体(goroutine)间的通信设施
//4)接口（interface）：对一组满足某个契约的类型的抽象

//3:结构体
type Rect struct {
	x, y          float64
	width, height float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r *Rect) Area2() float64 {
	return r.width * r.height
}

//在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以
//NewXXX来命名，表示“构造函数”
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//4:匿名组合
//在派生类中没有改写基类的成员方法时，相应的方法就被继承。例如下面的例子中
//调用foo.Foo()和调用foo.Base.Foo()效果一致

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("Base::Foo()")
}

func (base *Base) Bar() {
	fmt.Println("Base::Bar()")
}

type Foo struct {
	Base //注意这里是struct名称
}

func (foo *Foo) Bar() {
	fmt.Println("Foo:Bar()")
	foo.Base.Bar()
}

//另外，在Go语言中，你还可以以指针方式从一个类型派生：
//type Foo2 struct{
//	*Base
//}
//上面这段Go代码仍然有派生效果，只是Foo创建实例的时候，需要外部提供一个Base实例的指针

type X struct {
	Name string
}
type Y struct {
	X
	Name string
}

func main() {

	var a Integer = 1

	result := a.Less(2)

	fmt.Println("result = ", result)

	a.Add(20)
	fmt.Println("a = ", a)

	var c = [3]int{1, 2, 3}
	var d = c
	d[1]++
	fmt.Println("c:", c, " d:", d)

	var e = &c
	e[1]++
	fmt.Println("c:", c)

	var rect Rect = Rect{0, 0, 20, 30}
	var area float64
	area = rect.Area()
	fmt.Println("area:", area)

	area = 0
	area = rect.Area2()
	fmt.Println("area:", area)

	//结构体初始化
	rect1 := new(Rect)
	fmt.Println("(", rect1.x, ",", rect1.y, ",", rect1.width, ",", rect1.height, ")")

	rect2 := &Rect{}
	fmt.Println("(", rect2.x, ",", rect2.y, ",", rect2.width, ",", rect2.height, ")")

	rect3 := &Rect{0, 0, 200, 200}
	fmt.Println("(", rect3.x, ",", rect3.y, ",", rect3.width, ",", rect3.height, ")")

	rect4 := &Rect{width: 100, height: 200}
	fmt.Println("(", rect4.x, ",", rect4.y, ",", rect4.width, ",", rect4.height, ")")

	foo := &Foo{}
	foo.Bar()
	foo.Foo()
	foo.Base.Foo()

	var y Y = Y{X{"X"}, "female"}
	fmt.Println("name:", y.Name)
	fmt.Println("base name:", y.X.Name)

}
