package main

import (
	"bytes"
	"encoding/binary"
	"net"
)

/**
 * @Author: Lee
 * @Date: 2021/4/14 11:17
 * @Desc: tcp
 */

func main()  {
	listener,err := net.Listen("tcp","0.0.0.0:10000")
	if err != nil {
		return
	}
	for{
		// 建立socket连接
		conn,err := listener.Accept()
		if err != nil {
			continue
		}
		// 业务处理逻辑
		go process(conn)
	}
}

func process(conn net.Conn)  {
	msg := "hello"
	if msg,err:=pack(msg);err==nil{
		conn.Write(msg)
	}
}

//消息打包
func pack(message string) ([]byte, error) {
	// 读取消息的长度
	length := int32(len(message))
	pkg := new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.BigEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}