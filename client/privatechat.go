package client

import "fmt"

// 私聊模式
func (client *Client) PrivateChat() {
	// 查询现在有哪些人在线，并由接收线程自动显示
	client.conn.Write([]byte("whos" + "\n"))
	fmt.Println(">>>>选择您要私聊的对象: exit表示退出")
	var remoteName string
	fmt.Scanln(&remoteName)
	if remoteName == "exit" || remoteName == "" {
		return
	}
	fmt.Println(">>>>请输入您的消息: exit表示退出")
	var msg string
	fmt.Scanln(&msg)
	for remoteName != "exit" {
		for msg != "exit" {
			if msg != "" {
				_, err := client.conn.Write([]byte("to|" + remoteName + "|" + msg + "\n"))
				if err != nil {
					fmt.Println("消息发送失败")
					return
				}
				msg = ""
				fmt.Println(">>>>请输入您的消息: exit表示退出")
				fmt.Scanln(&msg)
			}
		}
		remoteName = ""
		fmt.Println(">>>>选择您要私聊的对象: exit表示退出")
		fmt.Scanln(&remoteName)
	}

}
