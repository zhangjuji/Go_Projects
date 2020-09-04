package processes

import (
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/utils"
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

		fmt.Println("mes = ", mes)
	}
}
