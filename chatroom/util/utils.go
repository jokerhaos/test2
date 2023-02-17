package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	common "test/chatroom/common/message"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) ReadPkg() (msg common.Message, err error) {
	buf := make([]byte, 8096)
	// 第一次读取，堵塞在这里读取长度消息
	_, err = this.Conn.Read(buf[:4])
	if err != nil {
		// err = errors.New("read pkg header error")
		return
	}
	// fmt.Println(buf[:n])
	// 根据buf[:4]转成一个 uint32类型
	pkgLen := binary.BigEndian.Uint32(buf)
	// 第二次读取，堵塞在这里读取消息内容，长度拿上面的
	n, err := this.Conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// err = errors.New("read pkg body error")
		return
	}

	err = json.Unmarshal(buf[:pkgLen], &msg)
	if err != nil {
		// err = errors.New("json decode err error")
		return
	}
	return
}

func (this *Transfer) SendMessage(msg common.Message) (err error) {
	sendData, _ := json.Marshal(msg)
	// 1.先发送长度，防止丢包
	pkgLen := uint32(len(sendData))
	fmt.Println("发送长度:", pkgLen)
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	_, err = this.Conn.Write(buf[:4])
	if err != nil {
		return err
	}
	// 2.发送消息数据
	_, err = this.Conn.Write(sendData)
	return
}
