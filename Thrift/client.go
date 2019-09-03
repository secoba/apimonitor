package Thrift

// package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "github.com/xuchengzhi/apimonitor/thriftserver/gen-go/demo/rpc"
	"github.com/xuchengzhi/apimonitor/thriftserver/gen-go/testone/rpc"
	"log"
	"net"
	"os"
	// "time"
)

// var thriftip = "192.168.248.188"

func ActRun(msg, ips string) string {
	//执行购买字体程序，msg为手机厂商如:huawei,ips 手机ip地址
	log.Println("actrun")
	// transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19090"))
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
		return "AutoPay服务未启动"
	}
	defer transport.Close()

	// for i := 0; i < 1000; i++ {
	// 	paramMap := make(map[string]string)
	// 	paramMap["name"] = "qinerg"
	// 	paramMap["passwd"] = "123456"
	// 	r1, e1 := client.FunCall(currentTimeMillis(), "login", paramMap)
	// 	fmt.Println(i, "Call->", r1, e1)
	// }
	// r1, e1 := client.TestOne(msg, ips)
	r1, e1 := client.TestOne(msg, ips)
	fmt.Println(r1, e1)
	if e1 == nil {
		return r1
	} else {
		fmt.Println(e1.Error())
		return "AutoPay服务异常"
	}
}

func ADB() string {
	//通过adb获取在线设备
	log.Println("adb")
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19091"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
		return "ADB服务未启动"
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:19090", " ", err)
		// os.Exit(1)
		return "ADB服务未启动"
	}
	defer transport.Close()
	res, e1 := client.Adb()
	if e1 == nil {
		return res
	} else {
		fmt.Println(e1.Error())
		return "ADB服务异常"
	}

}

func Atx() string {
	//启动atx服务
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
		return "Atx服务未启动"
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:19090", " ", err)
		// os.Exit(1)
		return "Atx服务未启动"
	}
	defer transport.Close()
	res, e1 := client.Atx()
	if e1 == nil {
		log.Println(res)
		return res
	} else {
		fmt.Println(e1.Error())
		return "Atx服务异常"
	}
}

func Order(app string) (string, bool) {
	//查询订单列表
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
		return "Order服务未启动", false
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:19090", " ", err)
		// os.Exit(1)
		return "Order服务未启动", false
	}
	defer transport.Close()
	res, e1 := client.Order(app)
	if e1 == nil {
		return res, true
	} else {
		fmt.Println(e1.Error())
		return "Order服务异常", false
	}
}

func Device() string {
	//获取在线设备列表
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
		return "Device服务未启动"
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 192.168.248.188:19090", " ", err)
		// os.Exit(1)
		return "Device服务未启动"
	}
	defer transport.Close()
	res, e1 := client.DevIce()
	if e1 == nil {
		return res
	} else {
		fmt.Println(e1.Error())
		return "Device服务异常"
	}
}

func DevUp() string {
	//更新设备在线状态
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(thriftip, "19090"))
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
		return "DevUp服务未启动"
	}
	defer transport.Close()

	res, e1 := client.Dev_UP()
	if e1 == nil {
		return res
	} else {
		fmt.Println(e1.Error())
		return "DevUp服务异常"
	}
}
