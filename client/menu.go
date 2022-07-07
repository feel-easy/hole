package client

import "fmt"

func (client *Client) menu() bool {
	fmt.Println(`1.公聊模式
2.私聊模式
3.更改用户名
0.退出`)
	var flag int
	fmt.Scanln(&flag)
	if flag >= 0 && flag < 4 {
		client.flag = flag
		return true
	} else {
		fmt.Println("您的输入不合法")
		return false
	}
}
