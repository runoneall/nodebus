package ipc

import (
	"fmt"

	"go.nanomsg.org/mangos/v3/protocol"
	"go.nanomsg.org/mangos/v3/protocol/req"
	_ "go.nanomsg.org/mangos/v3/transport/all"
)

func Connect(
	ipcName string,
) protocol.Socket {
	filePath := getIPCFile(ipcName)

	sock, err := req.NewSocket()
	if err != nil {
		panic(fmt.Errorf("不能初始化连接: %v", err))
	}

	if err := sock.Dial(fmt.Sprintf("ipc://%s", filePath)); err != nil {
		panic(fmt.Errorf("不能连接到服务: %v", err))
	}

	return sock
}
