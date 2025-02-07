package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server start...") //服务器启动了
	listen, err := net.Listen("tcp", "192.168.1.40:8888")
	if err != err {
		fmt.Println("listen err=", err)
		return
	}
	//监听成功以后
	//循环等待等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != err {
			fmt.Println("Accept err=", err)
		} else {
			fmt.Printf("等待连接成功，con = %v\n,接收到的客户端信息：%v\n", conn, conn.RemoteAddr().String())
		}
	}

}

//先启动服务器端，再启动客户端进行访问
