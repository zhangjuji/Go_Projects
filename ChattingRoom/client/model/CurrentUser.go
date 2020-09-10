package model

import (
	"lessons/GitHub/Go_Projects/ChattingRoom/common/message"
	"net"
)

type CurrentUser struct {
	Conn net.Conn
	message.User
}
