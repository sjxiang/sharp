package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var (
	verbose bool 
	timeout int
)

const DefaultTimeout = 0

func init() {
	flag.IntVar(&timeout, "w", DefaultTimeout, "无法建立或闲置超时的连接，读秒后")
	flag.BoolVar(&verbose, "v", false, "提供更多详情")
	flag.Parse()  // 先筛一遍
}


func checkError(err error) {
	if err == nil {
		return 
	}

	if verbose {
		fmt.Fprint(os.Stderr, err)
	}

	os.Exit(1)
}


func main() {
	
	// 解析参数，之气筛过 1 遍，还剩下 [host, port]
	args := flag.Args()
	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	host := os.Args[0]
	port := os.Args[1]


	// 连接服务器
	timeout := time.Duration(timeout) * time.Second
	conn, err := net.DialTimeout("tcp", host + ":" + port, timeout)
	checkError(err)
	defer conn.Close()  // 发送 FIN 包，指示目标端口断开连接

	go func ()  {
		// 读取标准输入数据，发送给目标端口
		io.Copy(conn, os.Stdin)	
	}()
	
	// 读取目标端口的数据，并写入到标准输出
	io.Copy(os.Stdout, conn)
}