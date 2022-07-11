/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"flag"
	"fmt"

	"github.com/feel-easy/hole/client"
)

var (
	serverIP   string
	serverPort string
)

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器的IP地址")
	flag.StringVar(&serverPort, "port", "8888", "设置服务器的连接端口号")
}

func main() {
	flag.Parse()
	client := client.GetClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>>>>>>服务器连接失败>>>>>>>>>")
		return
	}

	fmt.Println(">>>>>>>>>>服务器连接成功>>>>>>>>>>>")
	go client.DealResponse()
	client.Run()

	select {}
}
