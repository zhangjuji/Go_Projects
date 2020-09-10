package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/client/utils"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
)

type SmsProcess struct {
}

// 发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {
	// 1.创建一个 Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content

	smsMes.UserId = CurUser.UserId         // Id
	smsMes.UserStatus = CurUser.UserStatus // 状态

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("群发消息序列化失败 err = ", err)
		return
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("群发消息再次序列化失败 err = ", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("群发消息失败 err = ", err)
		return
	}

	return
}
