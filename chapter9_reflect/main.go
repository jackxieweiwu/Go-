// chapter9_reflect project main.go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4

	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())

}
