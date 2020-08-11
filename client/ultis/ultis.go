package ultis

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-code/pro/HaiLiang/commen/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时的缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	//conn只有在没有被关闭的情况下，才会阻塞
	//如果客户端关闭了conn，就不会阻塞，然后读不到东西就会一直报错
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("readPkg head errror")
		return
	}
	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("readPkg bogy errror")
		return
	}
	//把pkgLen反序列化成-->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //&一定要加上，不然mes就是空的
	if err != nil {
		fmt.Println("json.UnmarshaL err = ", err)
		return
	}
	return

}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度给对方
	var pkglen uint32
	pkglen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkglen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write err= ", err)
		return
	}

	//发送数据本身
	n, err = this.Conn.Write(data)
	if n != int(pkglen) || err != nil {
		fmt.Println("conn.Write err= ", err)
		return
	}
	return
}
