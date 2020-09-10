package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn   // 链接
	Buf  [8096]byte // 这是传输时，使用缓冲
}

// 编写一个 ReadPkg 读取函数
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	// buf := make([]byte, 8096)

	fmt.Println("读取客户端发送的数据......")

	_, err = this.Conn.Read(this.Buf[:4])

	if err != nil {
		fmt.Println("this.Conn.Read(this.Buf[:4]) err = ", err)
		return
	}

	// fmt.Println("读取到的长度 = ", this.Buf[:4])

	// 根据buf[:4] 转成一个	uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// 根据 pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("this.Conn.Read(this.Buf[:pkgLen]) err = ", err)
		return
	}

	// 把 pkgLen 反序列化成 message.Message
	// 技术就是一层窗户纸

	err = json.Unmarshal(this.Buf[:pkgLen], &mes) // 要传&mes 不然是空值
	if err != nil {
		fmt.Println("json.Unmarshal(Buf[:pkgLen], &mes) err = ", err)
		return
	}

	return
}

// 编写一个函数 WritePkg
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("this.Conn.Write(pkgBytes[0:4]) err = ", err)
		return
	}

	// 发送 data 本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("this.Conn.Write(data) err = ", err)
		return
	}
	return
}
