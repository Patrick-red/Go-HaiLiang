package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegistMesType           = "RegistMes"
	RegistResMesType        = "RegistResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType              = "SmsMes"
)

//定义状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code    int    `json:"code"`
	UserIds []int  //保存用户id的切片
	Error   string `json:"error"`
}

type RegistMes struct {
	User User `json:"user"`
}
type RegistResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	User           //匿名结构体
	Content string `json:"content"`
}
