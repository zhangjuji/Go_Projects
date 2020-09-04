package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据......")

	_, err = conn.Read(buf[:4])

	if err != nil {
		fmt.Println("conn.Read(buf[:4]) err = ", err)
		return
	}

	fmt.Println("读取到的长度 = ", buf[:4])

	// 根据buf[:4] 转成一个	uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根据 pkgLen 读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read(buf[:pkgLen]) err = ", err)
		return
	}

	// 把 pkgLen 反序列化成 message.Message
	// 技术就是一层窗户纸

	err = json.Unmarshal(buf[:pkgLen], &mes) // 要传&mes 不然是空值
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err = ", err)
		return
	}

	return
}

// 编写一个函数 writePkg
func writePkg(conn net.Conn, data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)

	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(pkgBytes[0:4]) err = ", err)
		return
	}

	// 发送 data 本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err = ", err)
		return
	}
	return
}
