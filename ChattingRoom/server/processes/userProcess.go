package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn // 链接
}

// 编写一个函数 ServerProcessLogin 函数，专门处理登陆请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心代码......
	// 1.先从 mes 中取出 mes.Data,并直接反序列化成 LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err = ", err)
		return
	}

	// 声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginMesType

	// 2.声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 验证先写死
	if loginMes.UserId == 100 && loginMes.Password == "123456" {
		// 合法
		loginResMes.Code = 200

	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用......"
	}

	// 3.将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err = ", err)
		return
	}

	// 4.将 data 赋值给 resMes
	resMes.Data = string(data)

	// 5.对 resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err = ", err)
		return
	}

	// 6.发送 data，我们将其封装到 writePkg 函数
	// 因为使用分层模式(mvc)，我们先创建一个Transfer实例
	t := &utils.Transfer{
		Conn: this.Conn,
	}
	err = t.WritePkg(data)

	return
}