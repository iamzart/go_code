package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("client start...")
	conn, err := net.Dial("tcp", "192.168.1.40:8888")
	if err != nil {
		fmt.Println("client start err,exit!", err)
		return
	}
	fmt.Println("client start success,conn=", conn)

}
