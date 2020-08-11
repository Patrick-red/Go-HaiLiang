package process

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
)

func outputGroupMes(mes *message.Message) {
	//反序列化
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
