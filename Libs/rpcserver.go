package rpcserver

import (
	"errors"
	"fmt"
	// "net/http"
	// "net/rpc"
	pb "apimonitor/Libs/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	fmt.Println("ddddd")
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Testss(args *Args, res *int) error {
	*res = args.A + args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// func main() {

// 	arith := new(Arith)
// 	rpc.Register(arith)
// 	rpc.HandleHTTP()
// 	fmt.Println("127.0.0.1:1234")
// 	err := http.ListenAndServe(":1234", nil)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// }

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
