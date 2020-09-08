package model

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 在服务器启动后，就初始化一个 UserDao 实例，把它做成全局变量，在需要和redis操作时，直接使用
// 定义一个 UserDao 结构体
// 完成对 User 的操作

type UserDao struct {
	pool *redis.Pool
}

var (
	MyUserDao *UserDao
)

// 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 绑定 UserDao 方法
// 根据用户 id 找用户
func (this *UserDao) getUserByID(conn redis.Conn, id int) (user *User, err error) {

	// 通过给定 id 去 redis 查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 在 users 哈希中，没有找到对应 id
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	// 反序列化 res 成 User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), user) err = ", err)
		return
	}

	return
}

// 登陆校验
// 1.完成对用户的验证
// 2.如果用户的 id 和 pwd 都正确，则返回一个 User 实例
// 3.如果用户的 id 或 pwd 都错误，则返回错误信息
func (this *UserDao) Login(userID int, password string) (user *User, err error) {

	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserByID(conn, userID)
	if err != nil {
		return
	}

	// 校验密码
	if user.Password != password {
		err = ERROR_USER_PWD // 密码错误
		return
	}

	return
}
