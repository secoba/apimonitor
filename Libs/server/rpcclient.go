package client

import (
	// "fmt"
	"log"
	// "net/rpc"
	pb "apimonitor/Libs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Rpcstr struct {
	Name string
}

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func Run() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: ", os.Args[0], "server")
	// 	os.Exit(1)
	// }
	// serverAddress := "127.0.0.1" //os.Args[1]

	// client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	// if err != nil {
	// 	log.Fatal("dialing:", err)
	// }
	// Synchronous call
	// par := make(map[string]string)
	// var res string
	// par["name"] = "haha"
	// err = client.Call("Say.Hello", par, &res)
	// if err != nil {
	// 	log.Fatal("arith error:", err)
	// }
	// fmt.Printf("Arith: %d\n", reply)

	// var quot Quotient
	// var args = Args{24, 56}
	// err = client.Call("Arith.Divide", args, &quot)
	// if err != nil {
	// 	log.Fatal("arith error:", err)
	// }
	// fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
