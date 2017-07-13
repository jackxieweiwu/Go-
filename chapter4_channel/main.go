// chapter4_channel project main.go
package main

import (
	"fmt"
)

//Go语言提供的另一种通信模型，即以消息机制而非共享内存作为通信方式
//消息机制认为每个并发单元是自包含的，独立的个体，并且都有自己的变量，但在不同并发单元间这些
//变量不共享。每个并发单元的输入和输出只有一种，那就是消息。这有点类似与进程的概念，每个进程
//不会被其他进程打扰，它只要做好自己的工作就可以了。不同进程间靠消息来通信，它们不会共享内存。

//Go语言提供的消息通信机制被称为channel
//“不要通过共享内存来通信，而应该通过通信来共享内存”

//channel是Go语言在语言级别提供的goroutine间的通信方式。我们可以使用channel在两个或者多个
//goroutine之间传递消息。channel是进程内的通信方式，因此通过channel传递对象的过程和调用
//函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，建议用分布式系统的方法
//来解决，比如使用Socket或者HTTP等通信。

//channel是类型相关的。也就是说，一个channel只能传递一种类型。这个类型需要在声明channel时
//指定。

func Count(ch chan int, number int) {
	fmt.Println("Counting(start):", number)
	ch <- number
	fmt.Println("Counting(end):", number)
}

func main() {
	chs := make([]chan int, 1)

	for i := 0; i < 1; i++ {
		chs[i] = make(chan int)

		go Count(chs[i], i+1)
	}

	for _, ch := range chs {
		number := <-ch
		fmt.Println("number:", number)
	}
}

//在上面的例子中，我们定义了一个包含10个channel的数组（名为chs)，并把数组中的每个channel分配给10个
//不同的goroutine。在每个goroutine的fmt.Println()函数完成后，我们通过ch <- 1语句向对应的channel
//中写入一个数据。在这个channel被读取之前，这个操作是阻塞的。在所有goroutine启动完成之后，我们通过<-ch
//语句从10个channel中依次读取数据。在对应的channel写入数据前，这个操作也是阻塞的。这样，就用channel
//实现了类似锁的功能，进而保证了goroutine完成后主函数才返回。

//1:基本语法
//var chanName chan ElementType
//与一般的变量声明不同的地方仅仅是在类型之前加了chan关键字。Element指定这个channel所能传递的元素类型
//例如：
// var ch chan int
// var chmap map[string]chan bool

//在channel的用法中，最常见的包括写入和读出。将一个数据写入（发送）至channel的语法很直观，如下：
// ch <- value

//向channel写入数据通常会导致程序阻塞，直到有其他goroutine从这个channel中读取数据。从channel
//中读取数据的语法是：
// value := <-ch
//如果channel之前没有写入数据，那么从channel中读取数据也会导致程序阻塞，直到channel中被写入数据为止。
