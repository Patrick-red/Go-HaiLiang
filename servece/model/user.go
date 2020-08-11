package model

//定义一个用户的结构体

type User struct {
	//为了序列化和反序列化成功，必须保证
	//用户信息的json字符串和结构体的字段的tag保持一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
