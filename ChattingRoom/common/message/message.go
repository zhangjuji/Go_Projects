package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterMesType         = "RegisterMes"
	RegisterResMesType      = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

// 用户状态常量
const (
	UserOnLine = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` // 消息类型
	Data string `json:"data"` // 消息的内容
}

// 定义两个消息..后面需要在增加
type LoginMes struct {
	UserId   int    `json:"userid"`   // 用户id
	Password string `json:"password"` // 密码
	UserName string `json:"username"` // 用户名
}

type LoginResMes struct {
	Code    int    `json:"code"`    // 返回状态码 500 表示该用户未注册 200 表示登陆成功
	UserIds []int  `json:"userids"` // 保存用户 id 的切片
	Error   string `json:"error"`   // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` // 类型就是 User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回码 400 表示该用户已经占用  200 表示注册成功
	Error string `json:"error"` // 返回错误信息
}

// 为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` // 用户id
	Status int `json:"status"` // 用户状态
}
