package main

import (
	"fmt"
	"net"
)

// // 读取函数
// func readPkg(conn net.Conn) (mes message.Message, err error) {

// 	buf := make([]byte, 8096)

// 	fmt.Println("读取客户端发送的数据......")

// 	_, err = conn.Read(buf[:4])

// 	if err != nil {
// 		fmt.Println("conn.Read(buf[:4]) err = ", err)
// 		return
// 	}

// 	fmt.Println("读取到的长度 = ", buf[:4])

// 	// 根据buf[:4] 转成一个	uint32类型
// 	var pkgLen uint32
// 	pkgLen = binary.BigEndian.Uint32(buf[:4])

// 	// 根据 pkgLen 读取消息内容
// 	n, err := conn.Read(buf[:pkgLen])
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Read(buf[:pkgLen]) err = ", err)
// 		return
// 	}

// 	// 把 pkgLen 反序列化成 message.Message
// 	// 技术就是一层窗户纸

// 	err = json.Unmarshal(buf[:pkgLen], &mes) // 要传&mes 不然是空值
// 	if err != nil {
// 		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err = ", err)
// 		return
// 	}

// 	return
// }

// // 编写一个函数 writePkg
// func writePkg(conn net.Conn, data []byte) (err error) {
// 	// 先发送一个长度给对方
// 	var pkgLen uint32
// 	pkgLen = uint32(len(data))

// 	var buf [4]byte
// 	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

// 	n, err := conn.Write(buf[0:4])
// 	if n != 4 || err != nil {
// 		fmt.Println("conn.Write(pkgBytes[0:4]) err = ", err)
// 		return
// 	}

// 	// 发送 data 本身
// 	n, err = conn.Write(data)
// 	if n != int(pkgLen) || err != nil {
// 		fmt.Println("conn.Write(data) err = ", err)
// 		return
// 	}
// 	return
// }

// // 编写一个函数 serverProcessLogin 函数，专门处理登陆请求
// func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
// 	// 核心代码......
// 	// 1.先从 mes 中取出 mes.Data,并直接反序列化成 LoginMes
// 	var loginMes message.LoginMes
// 	err = json.Unmarshal([]byte(mes.Data), &loginMes)
// 	if err != nil {
// 		fmt.Println("json.Unmarshal fail err = ", err)
// 		return
// 	}

// 	// 声明一个 resMes
// 	var resMes message.Message
// 	resMes.Type = message.LoginMesType

// 	// 2.声明一个 LoginResMes
// 	var loginResMes message.LoginResMes

// 	// 验证先写死
// 	if loginMes.UserId == 100 && loginMes.Password == "123456" {
// 		// 合法
// 		loginResMes.Code = 200

// 	} else {
// 		loginResMes.Code = 500
// 		loginResMes.Error = "该用户不存在，请注册再使用......"
// 	}

// 	// 3.将 loginResMes 序列化
// 	data, err := json.Marshal(loginResMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(loginResMes) err = ", err)
// 		return
// 	}

// 	// 4.将 data 赋值给 resMes
// 	resMes.Data = string(data)

// 	// 5.对 resMes 进行序列化，准备发送
// 	data, err = json.Marshal(resMes)
// 	if err != nil {
// 		fmt.Println("json.Marshal(resMes) err = ", err)
// 		return
// 	}

// 	// 6.发送 data，我们将其封装到 writePkg 函数
// 	err = writePkg(conn, data)

// 	return
// }

// // 编写一个 ServerProcessMes 函数
// // 根据客户端发送消息种类不同，决定调用哪个函数来处理
// func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

// 	switch mes.Type {
// 	case message.LoginMesType:
// 		// 处理登陆
// 		serverProcessLogin(conn, mes)
// 	case message.RegisterMesType:
// 		// 处理注册
// 	default:
// 		fmt.Println("消息类型不存在，无法处理......")
// 	}
// 	return
// }

// 处理和客户端的通讯
func process(c net.Conn) {

	// 延时关闭
	defer c.Close()

	// 调用总控
	pr := &Processor{
		Conn: c,
	}
	err := pr.process()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err = ", err)
		return
	}

}

func main() {

	// 提示信息
	fmt.Println("服务器[新的结构]在8889端口监听......")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	// 一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept err = ", err)
		}

		defer conn.Close()
		// 一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
