package Thrift

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "github.com/xuchengzhi/apimonitor/thriftserver/gen-go/demo/rpc"
	"github.com/xuchengzhi/apimonitor/thriftserver/gen-go/ThriftServer/rpc"
	"log"
	"net"
	"os"
	// "time"
)

func ThriftJavaServer() string {
	//调用java
	// transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("192.168.248.249", "11000"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
		return "ThriftService服务未启动"
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:11000", " ", err)
		return "Java服务未启动"
	}
	defer transport.Close()
	res, e1 := client.JavaServer()
	if e1 == nil {
		return res
	} else {
		log.Println(e1.Error())
		return "ThriftService服务异常"
	}

}
