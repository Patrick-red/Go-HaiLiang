package process

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/client/ultis"
	"go-code/pro/HaiLiang/commen/message"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType
	//创建SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	//序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("sms json fail err = ", err.Error())
		return
	}
	mes.Data = string(data)
	//序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("sms json fail err = ", err.Error())
		return
	}
	//发送mes
	tf := &ultis.Transfer{
		Conn: CurUser.Conn,
	}
	//
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("sms json fail err = ", err.Error())
		return
	}
	return

}
