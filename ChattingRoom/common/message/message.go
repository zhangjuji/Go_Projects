package message

const (
	LoginMesType       = "LoginMes"
	LoginResMesType    = "LoginResMes"
	RegisterMesType    = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
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
	Code  int    `json:"code"`  // 返回状态码 500 表示该用户未注册 200 表示登陆成功
	Error string `json:"error"` // 返回错误信息
}

type RegisterMes struct {
	User User `json:"user"` // 类型就是 User结构体
}

type RegisterResMes struct {
	Code  int    `json:"code"`  // 返回码 400 表示该用户已经占用  200 表示注册成功
	Error string `json:"error"` // 返回错误信息
}
