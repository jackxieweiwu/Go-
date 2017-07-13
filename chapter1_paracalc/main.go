// chapter1_paracalc project main.go
package main

import (
	"fmt"
	"reflect"
)

func sum(values []int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum
}

func change(values []int) {
	fmt.Printf("values:%d\n", &values[0])
	fmt.Printf("values:%d\n", values)
	values = append(values, 6, 7)
	fmt.Printf("dd:%p\n", &values[0])
	fmt.Printf("values:%p\n", values)

	fmt.Println("%#+v", values)
	for i := 0; i < len(values); i++ {
		values[i] += 1
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultChan := make(chan int, 2)

	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)

	sum1, sum2 := <-resultChan, <-resultChan

	fmt.Println("Result:", sum1, sum2, sum1+sum2)

	p := make([][]byte, 5)

	fmt.Printf("p: %p %p %p %p %p\n", &p[0], &p[1], &p[2], &p[3], &p[4])

	s := "AABBCCDD-EEFF-2233-4455-012364578921"
	format := "%08x-%04x-%04x-%04x-%012x"

	if _, err := fmt.Sscanf(s, format, &p[0], &p[1], &p[2], &p[3], &p[4]); err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	var b []int = []int{1, 2, 3, 4}
	var c []int = b
	//c = append(c, 5, 6, 7, 8, 9, 10, 11, 12)

	var d [5]int = [5]int{1, 2, 3, 4, 5}
	e := d

	c[0] = 100
	fmt.Println("type:", reflect.TypeOf(b))
	fmt.Printf("%p %p\n", b, c)
	fmt.Printf("%p,%p\n", &b[0], &b[1])
	fmt.Printf("%d\n", b[0])
	fmt.Printf("xx:%p\n", &b)
	fmt.Printf("%p %p\n", &d, &e)

	var slice []int
	if slice == nil {
		fmt.Println("slice is nil")
	} else {
		fmt.Println("slice is not nil")
	}

}
