package client

import "fmt"

// 更新用户名的业务方法
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>>>>请输入用户名>>>>>>>>>>")
	fmt.Scanln(&client.Name)
	_, err := client.conn.Write([]byte("rename|" + client.Name + "\n"))
	if err != nil {
		fmt.Println("更新出错:", err)
		return false
	}
	return true
}
