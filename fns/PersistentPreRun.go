package fns

import (
	"bytes"
	"fmt"
	"nodebus/cli"
	"nodebus/configm"
	"nodebus/ipc"

	"github.com/spf13/cobra"
)

func PersistentPreRun(cmd *cobra.Command, args []string) {
	configManager := configm.GetManager()

	switch *cli.UseCfgCenter {

	case false:
		if err := configManager.LoadJSON(); err != nil {
			panic(fmt.Errorf("不能加载文件: %v", err))
		}

	case true:
		conn := ipc.Connect("cfgcenter")
		defer conn.Close()

		if err := conn.Send([]byte("fetch")); err != nil {
			panic(fmt.Errorf("不能发送请求: %v", err))
		}

		resp, err := conn.Recv()
		if err != nil {
			panic(fmt.Errorf("不能接收响应: %v", err))
		}

		if err := configManager.LoadJSONFromReader(bytes.NewReader(resp)); err != nil {
			panic(fmt.Errorf("不能反序列化配置: %v", err))
		}

	}
}
