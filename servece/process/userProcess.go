package process2

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	"go-code/pro/HaiLiang/servece/model"
	"go-code/pro/HaiLiang/servece/ultis"
	"net"
)

type UserProcess struct {
	//分析字段
	Conn net.Conn
	//增加字段表示该链接是哪个用户的
	UserId int
}

//通知所有在线用户的方法
func (this *UserProcess) Notifyothers(userId int) {
	//遍历其他在线用户，然后发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyOthers(userId)
	}
}

func (this *UserProcess) NotifyOthers(userId int) {
	//组装消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline
	//序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	mes.Data = string(data)
	//mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json err", err)
		return
	}
	//发送var
	tf := &ultis.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("transfer err= ", err)
		return
	}
}

func (this *UserProcess) ServeProRegist(mes *message.Message) (err error) {
	var registMes message.RegistMes
	err = json.Unmarshal([]byte(mes.Data), &registMes)
	if err != nil {
		fmt.Println("json.un fail err = ", err)
		return
	}
	//1.先声明一个ResMes
	var resMes message.Message
	resMes.Type = message.RegistResMesType
	var registResMes message.RegistResMes

	//现在我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao去验证
	err = model.MyUserDao.Regist(&registMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registResMes.Code = 505
			registResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registResMes.Code = 506
			registResMes.Error = "注册发生未知错误"
		}
	} else {
		registResMes.Code = 200
	}
	data, err := json.Marshal(registResMes)
	if err != nil {
		fmt.Println("json fail err= ", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json fail err= ", err)
		return
	}
	//6.发送data，将其封装到writePkg函数
	//因为分层，先创建一个Tansfer实例，然后读取
	tf := &ultis.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//编写一个函数serveceProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServeProLogin(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.un fail err = ", err)
		return
	}
	//1.先声明一个ResMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//2.再声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//现在我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_NOTEXIST {
			loginResMes.Code = 500 //表示该用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 404
			loginResMes.Error = "未知错误"
		}

		//这里先测试，再返回具体的错误信息
	} else {
		loginResMes.Code = 200

		//用户登录成功，于是把它放入到userMgr中
		//将登录成功的userId赋给this
		this.UserId = loginMes.UserId

		userMgr.AddOnlienUser(this)
		this.Notifyothers(loginMes.UserId)
		//将当前在线用户的id放入到loginResMes.UserIds里面
		//遍历userMgr.onlin
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
		fmt.Println(user, "登录成功")
	}

	//如果用户id=100，密码=123456，认为合法，否则不合法
	/*if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法

	} else {
		//不合法


	}*/
	//3.将LoginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json fail err= ", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json fail err= ", err)
		return
	}
	//6.发送data，将其封装到writePkg函数
	//因为分层，先创建一个Tansfer实例，然后读取
	tf := &ultis.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
