package client

import "fmt"

// 公聊模式
func (client *Client) PublicChat() {
	var msg string
	fmt.Println(">>>>请输入您的消息: exit表示退出")
	for msg != "exit" {
		if _, err := fmt.Scanln(&msg); err != nil {
			return
		}
		if msg != "" {
			_, err := client.conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("消息发送失败")
				return
			}
			msg = ""
		}
	}
}
