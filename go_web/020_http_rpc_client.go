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

	serverAddress := "127.0.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)

	if err != nil {
		log.Fatal("aritch error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

}
