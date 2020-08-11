package modle

import (
	"go-code/pro/HaiLiang/commen/message"
	"net"
)

//因为很多地方会使用到，所以做成全局的
type CurUser struct {
	Conn net.Conn
	message.User
}
