package main

import (
	"fmt"
	"io"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/processes"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/utils"
	"net"
)

// 先创建一个Processor 的结构体
type Processor struct {
	Conn net.Conn
}

// 编写一个 ServerProcessMes 函数
// 根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	// 是否能接收到客户端发送的群发的消息
	fmt.Println("mes = ", mes)
	switch mes.Type {
	case message.LoginMesType:
		// 处理登陆
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		// 处理注册
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在，无法处理......")
	}
	return
}

func (this *Processor) process() (err error) {
	// 向客户端发送信息

	for {

		// 创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端退出......")
				return err
			} else {
				fmt.Println("readPkg(c) err = ", err)
				return err
			}
		}

		// fmt.Println("mes = ", mes)

		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
