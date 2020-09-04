package main

import (
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/processes"
	"os"
)

// 定义两个变量，一个表示用户的id，一个表示用户的密码
var userId int
var password string

func main() {

	// 接受用户选择
	var key int
	// 判断是否还继续显示菜单
	// var loop = true

	for true {
		fmt.Println("----------------------欢迎登陆多人聊天系统----------------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 4 请选择（1-3）：")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			// 说明用户要登陆
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &password)
			// loop = false
			up := &processes.UserProcess{}
			up.Login(userId, password)

		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("欢迎下次使用")
			os.Exit(0)
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

	}

	// // 更改用户的输入，显示新的提示信息
	// if key == 1 {

	// 	// 因为使用了新的程序结构，我们创建

	// 	// 先把登陆函数先到另外一个文件
	// 	login(userId, password)
	// 	// if err != nil {
	// 	// 	fmt.Println("登陆失败！")
	// 	// } else {
	// 	// 	fmt.Println("登陆成功！")
	// 	// }

	// } else if key == 2 {

	// 	fmt.Println("进行用户注册的逻辑...")
	// }
}
