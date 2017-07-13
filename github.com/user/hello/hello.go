package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
