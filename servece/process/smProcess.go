package process2

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/client/ultis"
	"go-code/pro/HaiLiang/commen/message"
	"net"
)

type SmsProcess struct {
	//...
}

//转发方法

func (this *SmsProcess) SendGroupMes(mes *message.Message) {
	//遍历服务器端的在线用户将消息转发出去
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("smProcess json err =  ", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("fail")
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendToEach(data, up.Conn)
	}
}
func (this *SmsProcess) SendToEach(data []byte, conn net.Conn) {
	tf := &ultis.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("fail err = ", err)
		return
	}
}
