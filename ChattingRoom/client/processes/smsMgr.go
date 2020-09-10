package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
)

func outputGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("客户端输出群发消息反序列化失败 err = ", err.Error())
		return
	}

	// 显示信息
	info := fmt.Sprintf("用户 id:\t%d 群发消息:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
