package process

import (
	"fmt"
	_ "fmt"
	"go-code/pro/HaiLiang/client/modle"
	"go-code/pro/HaiLiang/commen/message"
)

//
var onlineUsers map[int]*message.User = make(map[int]*message.User, 100)
var CurUser modle.CurUser //用户登陆后，完成对CurUser的初始化

//
func outputOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {

		fmt.Println("用户id:\t", id)
	}
}

func updataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status

	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
