package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/utils"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
	"os"
)

type Server struct {
}

// 1.显示登陆成功界面......
func (this *Server) ShowMenu() {
	fmt.Println("---------恭喜登陆成功---------")
	fmt.Println("-------1.显示在线用户列表-------")
	fmt.Println("-------2.发送消息-------")
	fmt.Println("-------3.信息列表-------")
	fmt.Println("-------4.退出系统-------")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表")

	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
		os.Exit(0)
	default:
		fmt.Println("您输入的内容有误")
	}

}

func (this *Server) ServerMesProcess(conn net.Conn) {
	// 创建一个 Transfer 实例，不停地读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {

		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err = ", err)
			return
		}

		switch mes.Type {
		case message.NotifyUserStatusMesType: // 有人上线了
			// 1.取出.NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			// 2.把这个用户的信息，状态保存到客户端的 map 中
			updateUserStatus(&notifyUserStatusMes)
		default:
			fmt.Println("服务器端返回了一个未知的消息类型")
		}
	}
}
