package ipc

import (
	"fmt"

	"go.nanomsg.org/mangos/v3/protocol/rep"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

func Serv(
	ipcName string,
	onRecv func(data []byte, responder func(data []byte)),
) {
	filePath := getIPCFile(ipcName)
	cleanIPCFile(filePath)

	sock, err := rep.NewSocket()
	if err != nil {
		panic(fmt.Errorf("不能初始化连接: %v", err))
	}

	if err := sock.Listen(fmt.Sprintf("ipc://%s", filePath)); err != nil {
		panic(fmt.Errorf("不能监听连接: %v", err))
	}

	fmt.Println("启动服务:", filePath)

	for {
		data, err := sock.Recv()
		if err != nil {
			fmt.Println("不能读取消息:", err)
			continue
		}

		responder := func(data []byte) {
			if err := sock.Send(data); err != nil {
				fmt.Println("不能发送响应:", err)
			}
		}

		go onRecv(data, responder)
	}
}
