// chapter4_1 project main.go
package main

import (
	"fmt"
)

func Add(x, y int) {
	result := x + y
	fmt.Println(result)
}
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
