package model

// 定义一个用户的结构体

type User struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"username"`
}
