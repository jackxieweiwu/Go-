// chapter6_hash2.go project main.go
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func main() {
	TestFile := "123.txt"
	infile, err := os.Open(TestFile)
	if err == nil {
		md5Inst := md5.New()
		io.Copy(md5Inst, infile)
		fmt.Printf("%x   %s\n", md5Inst.Sum([]byte("")), TestFile)

		sha1Inst := sha1.New()
		io.Copy(sha1Inst, infile)
		fmt.Printf("%x   %s\n", sha1Inst.Sum([]byte("")), TestFile)
	}
}
