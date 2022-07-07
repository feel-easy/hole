package client

import "fmt"

// 公聊模式
func (client *Client) PublicChat() {
	fmt.Println(">>>>请输入您的消息: exit表示退出")
	var msg string
	fmt.Scanln(&msg)
	for msg != "exit" {
		if msg != "" {
			_, err := client.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("消息发送失败")
				return
			}
			msg = ""
			fmt.Println(">>>>请输入您的消息: exit表示退出")
			fmt.Scanln(&msg)
		}

	}
}
