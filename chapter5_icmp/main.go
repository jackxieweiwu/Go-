package main

import (
	//"bytes"
	"fmt"
	//	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}

	service := os.Args[1]
	fmt.Println("service1:", service)

	conn, err := net.Dial("ip4:icmp", service)

	checkError(err)

	var msg [512]byte

	/*
	   	0                   1                   2                   3
	   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |     Type      |     Code      |          Checksum             |
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |                             unused                            |
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |      Internet Header + 64 bits of Original Data Datagram      |
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   *
	   *8、请求回显或回显应答(Echo or Echo Reply Message)
	   *
	   0                   1                   2                   3
	   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |     Type      |     Code      |          Checksum             |
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |           Identifier          |        Sequence Number        |
	   +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	   |     Data ...
	   +-+-+-+-+-
	   Type(8)是请求回显报文(Echo)；Type(0)是回显应答报文(Echo Reply)。
	   请求回显或回显应答报文属于查询报文。Ping就是用这种报文进行查询和回应。
	*/

	msg[0] = 8  //echo
	msg[1] = 0  //code 0
	msg[2] = 0  //checksum
	msg[3] = 0  //checksum
	msg[4] = 0  //identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  //sequence[0]
	msg[7] = 37 //sequence[1]

	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 0xFF)

	_, err = conn.Write(msg[0:len])
	checkError(err)
	fmt.Println("Write success...")

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Get response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Identifier matches")
	}

	os.Exit(0)

}

func checkSum(msg []byte) uint16 {
	sum := 0
	n := 0

	for n+1 < len(msg) {
		sum += (int(msg[n]) << 8) | int(msg[n+1])
		n++
	}

	if n < len(msg) {
		sum += (int(msg[n]) << 8)
	}

	sum = (sum >> 16) + (sum & 0xFFFF)
	sum += (sum >> 16)

	return uint16(^sum)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
