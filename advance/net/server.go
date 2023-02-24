package main

import (
	"io"
	"net"
)

/*
	TCP 服务器
*/

func Serve(addr string) error {
	// 开始监听端口
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		// 创建连接
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func ()  {
			// 处理请求
			handle(conn)
		}()
	}
}


func handle(conn net.Conn) {
	for {
		// 读取数据
		reqBs := make([]byte, 8)
		_, err := conn.Read(reqBs)
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			// 一般错误，不让它触发 Close()，懒得管
			_ = conn.Close()
			return 
		}
		if err != nil {
			continue
		}	

		// 处理数据
	
		
		// 写回响应
		respBs := []byte("") 
		_, err = conn.Write(respBs)	
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return 
		}
	}

}


/*

三次握手和四次挥手


在真实环境下，接收数据的长度时不确定的。

	1. 特殊字符隔开，但要考虑转义 e.g. '\n'
	2. 传递消息长度


优化 

	1. 处理连接时，开 goroutine，多路复用。
	2. 处理消息时，读写单开 goroutine，配合 channel 传递数据，TCP 双工啊！
	
*/
