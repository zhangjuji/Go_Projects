package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/utils"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	// 取出 mes 的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("服务器端群发消息反序列化失败 err = ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("mes 序列化失败 err = ", err)
		return
	}

	// 遍历服务器端的 map
	for id, up := range userMgr.onlineUsers {
		// 过滤掉自身
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}

	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("服务器端转发消息失败 err = ", err)
		return
	}
}
