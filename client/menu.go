package client

import "fmt"

func (client *Client) menu() bool {
	fmt.Println(`1.公聊模式
2.私聊模式
0.退出`)
	var flag int
	if _, err := fmt.Scanln(&flag); err != nil {
		fmt.Println("您的输入不合法")
		return false
	}
	if flag >= 0 && flag < 4 {
		client.flag = flag
		return true
	} else {
		fmt.Println("您的输入不合法")
		return false
	}
}
