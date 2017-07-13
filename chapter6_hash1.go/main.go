// chapter6_hash1.go project main.go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hi,pandaman!"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(TestString))
	Result = sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)
}
