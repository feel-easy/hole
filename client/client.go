package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 定义Client类型
type Client struct {
	IP   string
	Port string
	Name string
	conn net.Conn
	flag int
}

func GetClient(IP, Port string) *Client {
	conn, err := net.Dial("tcp", IP+":"+Port)
	if err != nil {
		fmt.Println("客户端连接失败，错误信息为:", err)
		return nil
	}
	return &Client{
		IP:   IP,
		Port: Port,
		conn: conn,
		flag: 999,
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
			// 死循环，反复执行menu函数
		}
		// Go的switch不需要break语句
		switch client.flag {
		case 1:
			// 公聊模式代码块
			fmt.Println("您已进入公聊模式")
			client.PublicChat()
			break
		case 2:
			// 私聊模式代码块
			fmt.Println("您已进入私聊模式")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("您选择了更改用户名")
			client.UpdateName()
			break
		}
	}
}

// 监听服务器返回消息的方法，单开go程
func (client *Client) DealResponse() {
	// 一种读入连接并显示到标准输出的简写方法
	fmt.Println("客户端显示方法执行")
	io.Copy(os.Stdout, client.conn)
}
