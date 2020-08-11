package process2

import (
	"fmt"
)

//因为UserMgr实例在服务器端只有一个
//很多地方会用，所以定义成一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlienUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DeleteOnlienUser(userId int) {
	delete(this.onlineUsers, userId)
}

//返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

//根据ID返回对应的值
func (this *UserMgr) GetOnlienUserById(userId int) (up *UserProcess, err error) {
	//带检测的从map取出一个值
	up, ok := this.onlineUsers[userId]
	if !ok { //说明查找的用户当前不在线
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
