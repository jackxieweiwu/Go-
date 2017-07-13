// chapter5_simplehttp project main.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

//下面我们建立TCP链接来实现初步的HTTP协议，通过向网络主机发送HTTP Head请求，读取网络主机
//返回的信息
func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	fmt.Println("service:", service)
	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("GET /index.html HTTP/1.0\r\n\r\n"))
	checkError(err)

	fmt.Println("Write success...")

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
