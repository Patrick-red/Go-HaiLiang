package model

import (
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	_ "net"

	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后，就初始化一个userDao实例
//把他做成一个全局变量，在需要和redis操作时，就直接用
var (
	MyUserDao *UserDao
)

//定义一个UserDao的结构体
//完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//UserDao应该提供的方法
//1.根据用户id返回一个User实例和err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在users哈希中，没有找到对应的id
			err = ERROR_USER_NOTEXIST
		}
		return
	}
	//这里要把res反序列化成User实例

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	return
}

//完成登录的校验
//1.Login 完成对用户的验证
//2.如果用户id和密码都正确，返回一个user实例
//3.如果用户id或密码有错，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从UserDao中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这是证明用户获取到了
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Regist(user *message.User) (err error) {
	//先从UserDao中取出一根链接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//说明id还没被注册过
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//开始入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("注册入库错误err = ", err)
		return
	}
	return
}
