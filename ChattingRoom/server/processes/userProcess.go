package processes

import (
	"encoding/json"
	"fmt"
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/model"
	"lessons/GitHub/Go_Projects/ChattingRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn   net.Conn // 链接
	UserId int
}

func (this *UserProcess) NotifyOtherOnlineUser(userId int) {

	// 遍历 onlineUsers，然后一个一个地发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		// 过滤自己
		if id == userId {
			continue
		}
		up.NotifyMeOnline(userId)

	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {

	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnLine

	// 序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes 序列化出错，err = ", err)
		return
	}

	// 将序列化的 notifyUserStatusMes 赋值给 mes.Data
	mes.Data = string(data)

	// 对 mes 再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes 序列化出错，err = ", err)
		return
	}

	// 创建 Transfer 实例，发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("上线通知他人出错，err = ", err)
		return
	}

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {

	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("registerMes json.Unmarshal err = ", err)
		return
	}

	// 声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	// 3.将 registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) err = ", err)
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

	// 到 redis 数据库去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.Password)

	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		// 这里用户登陆成功，我们就把该登陆成功的用户放入到 UserMgr 中
		// 将登陆成功的用户的 id 赋给 this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		// 通知其他的在线用户，我上线了
		this.NotifyOtherOnlineUser(loginMes.UserId)
		// 将当前在线用户的 id 放入到 loginResMes.UserIds
		// 遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		fmt.Println(user, "登陆成功")
	}

	// // 验证先写死
	// if loginMes.UserId == 100 && loginMes.Password == "123456" {
	// 	// 合法
	// 	loginResMes.Code = 200

	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册再使用......"
	// }

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
