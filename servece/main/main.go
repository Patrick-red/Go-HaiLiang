package main

import (
	"fmt"
	"go-code/pro/HaiLiang/servece/duoyu"
	"go-code/pro/HaiLiang/servece/model"
	"net"
	"time"
)

/* func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	//conn只有在没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，就不会阻塞，然后读不到东西就会一直报错
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("readPkg head errror")
		return
	}
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("readPkg bogy errror")
		return
	}
	//把pkgLen反序列化成-->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes) //&一定要加上，不然mes就是空的
	if err != nil {
		fmt.Println("json.UnmarshaL err = ", err)
		return
	}
	return

}

func writePkg(conn net.Conn, data []byte) (err error) {
	//先发送长度给对方
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

	//发送数据本身
	n, err = conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Write err= ", err)
		return
	}
	return
} */
/*
//编写一个函数serveceProcessLogin函数，专门处理登录请求
func serveProLogin(conn net.Conn, mes *message.Message) (err error) {
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
	//如果用户id=100，密码=123456，认为合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200
	} else {
		//不合法
		loginResMes.Code = 500 //表示该用户不存在
		loginResMes.Error = "该用户不存在"
	}
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
	err = writePkg(conn, data)
	return
}

//编写一个ServiceProMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serveproMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		err = serveProLogin(conn, mes)
	case message.RegistMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
*/
//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//这里调用总控,创建一个
	processor := &duoyu.Processor{
		Conn: conn,
	}
	err := processor.Process2()
	if err != nil {
		fmt.Println("协程出错 err = ", err)
		return
	}
}

//这里对userDao初始化
func initUserDao() {
	//这里的pool本身就是一个全局变量
	//这里需要注意初始化顺序问题1
	//initpool在initUserDao
	model.MyUserDao = model.NewUserDao(duoyu.Pool)
}

func main() {
	//当服务器启动时，我们就去初始化我们的redis的链接池
	duoyu.InitPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()

	//提示信息
	fmt.Println("新的服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.listen err = ", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来连接服务器。。。。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯...
		go process(conn)
	}

}
