package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

/**
 * @Author: Lee
 * @Date: 2021/4/14 11:17
 * @Desc:
 */

const (
	BUFF_SIZE = 1024
	HEAD_LEN = 4
)
func main()  {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:10000",time.Second*3)
	if err!=nil {
		panic(err)
	}
	defer conn.Close()
	buffer := bytes.NewBuffer(make([]byte, 0, BUFF_SIZE))
	for {
		rb := make([]byte,BUFF_SIZE)
		rn,err := conn.Read(rb)
		if err!=nil||err == io.EOF{
			break
		}
		buffer.Write(rb[0:rn])
		for {
			if buffer.Len()>=HEAD_LEN{
				head := make([]byte, HEAD_LEN)
				_, err = buffer.Read(head)
				if err != nil {
					return
				}
				bodyLen := BytesToInt(head)
				if buffer.Len() >= bodyLen {
					body := make([]byte, bodyLen)
					_, err = buffer.Read(body[:bodyLen])
					if err != nil {
						return
					}
					fmt.Println(string(body))
				}else {
					break
				}
			}else {
				break
			}
		}
	}
}
//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}