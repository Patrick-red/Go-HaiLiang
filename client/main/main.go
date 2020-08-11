package main

import (
	"fmt"
	"go-code/pro/HaiLiang/client/process"
	"os"
)

//定义用户的id和密码
var userId int
var userPwd string
var userName string

func main() {
	var key int //用来接受用户选择
	//var loop = true //用来判断是否继续显示菜单
	for {
		fmt.Println("--------------欢迎登陆多人聊天系统---------------")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择（1-3）")
		fmt.Scanf("%d\n", &key) //\n一定要
		switch key {            //key要是整数，不能是字符串
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入你的id")     //说明用户要登陆
			fmt.Scanf("%d\n", &userId) //\n一定要，不然会以为回车是下面的userPwd，直接运行完了
			fmt.Println("请输入密码")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入姓名：")
			fmt.Scanf("%s\n", &userName)
			//调用UserProcess完成注册
			up := &process.UserProcess{}
			up.Regsiter(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
			//loop = false
		default:
			fmt.Println("输入有误，请重新输入")

		}

	}
	/*if key == 1 {

		//先把登陆的函数写到另外的文件
		//err := login(userId, userPwd)
		//if err != nil {
		//	fmt.Println("登陆失败")
		//} else {
		//	fmt.Println("登陆成功")
		//}
	} else if key == 2 {
		fmt.Println("用户注册")

	}*/
}
