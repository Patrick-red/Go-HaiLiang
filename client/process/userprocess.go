package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	"go-code/pro/HaiLiang/servece/ultis"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要
}

//关联一个用户登录的方

//完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//fmt.Printf("userId = %d , userPwd = %s\n", userId, userPwd)
	//return nil
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	//延时关闭
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//4.将LoginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err = ", err)
		return
	}
	//5.把data赋给mes.Data
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ", err)
		return
	}
	//7.现在data就是要发送的消息
	//7.1先把data长度发送个服务器
	//先获取data长度，再转成一个表示长度的切片
	var pkglen uint32
	pkglen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err= ", err)
		return
	}
	fmt.Printf("客户端发送消息的长度成功%d , 内容是%s", len(data), string(data))
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err= ", err)
		return
	}
	//休眠20
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20.。")

	//这里还需要处理服务器返回的消息
	//创建一个Transfer实例
	tf := &ultis.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() //mes就是

	if err != nil {
		fmt.Println("readpkg fail err = ", err)
		return
	}
	//将mes的Data反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//显示当前在线用户列表
		//初始化
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("用户列表如下显示")
		for _, v := range loginResMes.UserIds {
			if v == userId {
				continue
			}
			fmt.Println("用户id= ", v)
			//完成onlineUsers初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user

		}
		fmt.Println("\n\n")
		//1.显示登录成功的菜单
		//这里启动协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

//注册方法
func (this *UserProcess) Regsiter(userId int, userPwd string,
	userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}
	//延时关闭
	defer conn.Close()

	var mes message.Message
	mes.Type = message.RegistMesType
	//3.创建一个LoginMes结构体
	var registerMes message.RegistMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err = ", err)
		return
	}
	//5.把data赋给mes.Data
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ", err)
		return
	}
	//这里还需要处理服务器返回的消息
	//创建一个Transfer实例
	tf := &ultis.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册write错误 err = ", err)
	}
	mes, err = tf.ReadPkg() //mes就是registerResMes

	if err != nil {
		fmt.Println("readpkg fail err = ", err)
		return
	}
	//将mes的Data反序列化成LoginResMes
	var registerResMes message.RegistResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}
