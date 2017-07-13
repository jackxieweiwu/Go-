// chapter4_select project main.go
package main

import (
	"fmt"
)

//这个程序实现了一个随机向ch中写入一个0或者1的过程
func main() {

	//注意：这里需要创建一个带1个缓冲的channel，假若采用不带缓冲的channel会造成死锁。
	ch := make(chan int, 1)

	for {
		select {
		case ch <- 0:
		case ch <- 2:
		}

		i := <-ch
		fmt.Println("Value received:", i)
	}

}

//2:缓冲机制
//要创建带缓冲的channel，其实也很容易:
// c := make(chan int,1024)
// 在调用make()时将缓冲区大小作为第二个参数传入即可，比如上面我们就创建了一个大小为1024的int类型
//channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被填完之前都不会阻塞。

//从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，但我们也可以使用range关键字
//来实现更为简便的循环读取：
//for i := range c{
//	fmt.Println("Received:",i)
//}

//3: 超时机制
//Go语言没有提供直接的超时处理机制，但我们可以利用select机制。虽然select机制不是专为超时而设计的，却能
//很方便地解决超时问题。因为select的特点是只要其中一个case已经完成，程序就会继续往下执行，而不会考虑其他
//case的情况
//基于此特性，我们来为channel实现超时机制：

//首先，我们实现并执行一个匿名的超时等待函数

//timeout := make(chan bool,1)

//go func (){
//	time.Sleep(1e9)			//等待一秒钟
//	timeout <- true
//}()

////然后我们把timeout这个channel利用起来
//select {
//	case <- ch:
//	//从ch中读取到数据
//	case <- timeout:
//	//一直没有从ch中读取到数据，但从timeout中读取到了数据
//}

//这样使用select机制可以避免永久等待的问题，因为程序会在timeout中获取到一个数据后继续执行。无论对ch的读取是否还处于等待
//状态，从而达到1秒超时的效果。

//这种写法看起来是一个小技巧，但却是在Go语言开发中避免channel通信超时的最有效方法。

//4:单项channel
//单项channel只是双向channel的一种使用限制而已。

//我们在将一个channel变量传递到一个函数时，可以将其指定为单向channel变量，从而限制该函数中可以对此channel的操作，比如
//只能往这个channel写，或者只能从这个channel读。

//单向channel变量的声明非常简单，如下：
//var ch1 chan int       //ch1是一个正常的channel，不是单向的
//var ch2 chan<- float64 //ch2是一个单向channel，只用于写float64数据
//var ch3 <-chan int     //ch3是单向channel，只用于读取int数据

//单向channel的初始化：
//ch4 := make(chan int)
//ch5 := <-chan int(ch4)  //ch5就是一个单向的读取channel
//ch6 := chan<- int(ch4)	//ch6就是一个单向的写入channel
////基于ch4，我们通过类型转换初始化了两个单向channel：单向读的ch5和单向写的ch6

//5:关闭channel
// close(ch)

//如何判断一个channel是否已经被关闭？我们可以在读取的时候使用多重返回值的方式：
//x,ok := <-ch
//此时，我们只需要看第二个bool返回值，如果返回值是false则表示ch已经关闭

//5:sync.Once实现全局唯一性操作
//var a string
//var once sync.Once

//func setup() {
//	a = "hello,world"
//}

//func doprint() {
//	once.Do(setup)
//	print(a)
//}

//func twoprint() {
//	go doprint()
//	go doprint()
//}
