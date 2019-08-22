package main

//client.go

import (
	"fmt"
	pd "github.com/xuchengzhi/apimonitor/Grpc/lib"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

const (
	address = "localhost:8082"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pd.NewFormatDataClient(conn)

	r, err := c.DoFormat(context.Background(), &pd.Actionrequest{Text: "test", Corpus: pd.Actionrequest_NEWS})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	fmt.Println(r.Age)
	fmt.Println(reflect.TypeOf(r.Result))
	for k, v := range r.Result {
		fmt.Println(k, v)
		fmt.Println(v.Snippets)
	}
}
