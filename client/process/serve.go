package process

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	"go-code/pro/HaiLiang/servece/ultis"
	"net"
	"os"
)

//显示登录成功后的界面...
func ShowMenu() {
	fmt.Println("----------恭喜登录成功---------")
	fmt.Println("----------1.显示在线用户列表---------")
	fmt.Println("----------2.发送消息---------")
	fmt.Println("----------3.信息列表---------")
	fmt.Println("----------4.退出系统---------")
	fmt.Println("----------请选择（1-4）---------")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("请输入内容：\n")
		fmt.Println("下来了")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("查看信息列表")
	case 4:
		fmt.Println("你退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入不对")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	//创建一个Transfer实例，不停的读取服务器发送的消息
	tf := &ultis.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.read err = ", err)
			return
		}
		//如果没错，又是下一步处理逻辑
		//fmt.Println("mes = ", mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1.取出NotifyUserStatusMes
			//2.把该用户状态保存到客户端map[int]user
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updataUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("未知消息类型")
		}
	}
}
