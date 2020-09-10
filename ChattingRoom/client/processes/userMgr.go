package processes

import (
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/model"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
)

// 客户端要维护的在线用户 map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 100)

// 全局变量
var CurUser model.CurrentUser // 用户登陆成功后，初始化

// 在客户端显示当前在线用户
func showOnlineUser() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

// 处理返回的 NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status

	onlineUsers[notifyUserStatusMes.UserId] = user

	showOnlineUser()
}
