package duoyu

import (
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	"go-code/pro/HaiLiang/servece/process"
	"go-code/pro/HaiLiang/servece/ultis"
	"io"
	"net"
)

//先创建一个Processor的结构体

type Processor struct {
	Conn net.Conn
}

//编写一个ServiceProMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServeproMes(mes *message.Message) (err error) {
	//是否能接受群发
	fmt.Println("mes=", mes.Type)
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServeProLogin(mes)
	case message.RegistMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServeProRegist(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) Process2() (err error) {
	//循环读取客户端发送的信息
	for {
		//创建以恶搞Transfer实例完成读包的工作
		tf := &ultis.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出。。。")
				return err
			} else {
				fmt.Println("readPkg fial err = ", err)
				return err
			}
		}
		//fmt.Println("mes =", mes)
		err = this.ServeproMes(&mes)
		if err != nil {
			fmt.Println("serveproMes err = ", err)
			return err
		}

	}
}
