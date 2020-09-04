package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
)

// 写一个函数，完成登陆
func login(userId int, password string) (err error) {

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

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.Password = password

	// 4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes json.Marshal err = ", err)
		return
	}

	// 5.把data赋给 mes.Data字段
	mes.Data = string(data)

	// 6.将mes进行序列化
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

	// fmt.Printf("客户端发送消息长度为 %d,内容为 %s\n", len(data), string(data))
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err = ", err)
		return
	}

	// 反序列化成 mes.Data
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
