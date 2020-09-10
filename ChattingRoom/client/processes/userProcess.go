package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/utils"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, password string, username string) (err error) {
	// 1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3.创建一个 registerMes 结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.Password = password
	registerMes.User.UserName = username

	// 4.将 registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("registerMes json.Marshal err = ", err)
		return
	}

	// 5.把 data 赋给 mes.Data 字段
	mes.Data = string(data)

	// 6.将 mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal err = ", err)
		return
	}

	// 7.创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err = ", err)
	}

	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	// 反序列化成 mes.Data
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，可以重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return
}

// 给关联一个用户登录的方法
// 写一个函数，完成登陆
func (this *UserProcess) Login(userId int, password string) (err error) {

	// 开始订协议
	// fmt.Printf("userId = %d password = %s\n", userId, password)
	// return nil

	// 1.链接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 2.准备通过 conn 发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个 LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.Password = password

	// 4.将 loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes json.Marshal err = ", err)
		return
	}

	// 5.把 data 赋给 mes.Data 字段
	mes.Data = string(data)

	// 6.将 mes 进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marshal err = ", err)
		return
	}

	// 7.发送data给服务器
	// 7.1 先发送data的长度
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(pkgBytes[0:4]) err = ", err)
		return
	}

	fmt.Println("客户端发送消息长度成功")

	// 7.2 然后发送data消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}

	// 这里还需要处理服务器端返回的消息

	fmt.Printf("客户端发送消息长度为 %d,内容为 %s\n", len(data), string(data))
	// 创建一个 Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	// 反序列化成 mes.Data
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		// fmt.Println("登陆成功")

		// 显示当前在线用户列表，遍历 loginResMes.UserIds
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {

			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
		}
		fmt.Print("\n\n")

		// 客户端启动一个协程
		// 该协程保持和服务器端的通讯,如果服务器有数据推送给客户端
		// 则接受并显示在客户端的终端
		server := Server{}
		go server.ServerMesProcess(conn)

		// 1.循环显示菜单
		for {
			server.ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
