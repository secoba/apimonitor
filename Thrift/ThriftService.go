package Thrift

// package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "github.com/xuchengzhi/apimonitor/thriftserver/gen-go/demo/rpc"
	"github.com/xuchengzhi/apimonitor/thriftserver/gen-go/ThriftServer/rpc"
	"log"
	"net"
	"os"
	"time"
)

var thriftip = "192.168.248.188"

func ThriftAtxServer() string {
	//执行购买字体程序，msg为手机厂商如:huawei,ips 手机ip地址
	log.Println("actrun")
	// transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "11000"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:19090", " ", err)
		// os.Exit(1)
		return "ThriftService服务未启动"
	}
	defer transport.Close()

	// for i := 0; i < 1000; i++ {
	//  paramMap := make(map[string]string)
	//  paramMap["name"] = "qinerg"
	//  paramMap["passwd"] = "123456"
	//  r1, e1 := client.FunCall(currentTimeMillis(), "login", paramMap)
	//  fmt.Println(i, "Call->", r1, e1)
	// }
	// r1, e1 := client.TestOne(msg, ips)
	r1, e1 := client.AtxServer()
	fmt.Println(r1, e1)
	if e1 == nil {
		return r1
	} else {
		fmt.Println(e1.Error())
		return "ThriftService服务未启动"
	}
}

func ThriftAdbServer() string {
	//通过adb获取在线设备
	log.Println("adb")
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "11000"))
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
		// os.Exit(1)
		return "ThriftService服务未启动"
	}
	defer transport.Close()
	res, e1 := client.AdbServer()
	if e1 == nil {
		return res
	} else {
		fmt.Println(e1.Error())
		return "ThriftService服务异常"
	}

}

func ThriftActServer() string {
	//启动atx服务
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "11000"))
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
		// os.Exit(1)
		return "ThriftService服务未启动"
	}
	defer transport.Close()
	res, e1 := client.ActServer()
	if e1 == nil {
		return res
	} else {
		fmt.Println(e1.Error())
		return "ThriftService服务未启动"
	}
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
