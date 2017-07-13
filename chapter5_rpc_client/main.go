// chapter5_rpc_client project main.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:4321")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{7, 8}

	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

	args = &Args{5, 2}
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done

	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("Arith:%d / %d = %d %d\n", args.A, args.B, quotient.Quo, quotient.Rem)
}
