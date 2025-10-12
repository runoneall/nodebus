package fns

import (
	"encoding/json"
	"fmt"
	"nodebus/configm"
	"nodebus/ipc"

	"github.com/spf13/cobra"
)

func CfgCenterServer(cmd *cobra.Command, args []string) {
	configManager := configm.GetManager()

	ipc.Serv("cfgcenter", func(data []byte, responder func(data []byte)) {
		raise := func(err string) {
			responder([]byte(err))
		}

		command := string(data)
		switch command {

		case "fetch":
			jsonData, err := json.Marshal(configManager.ItemGetAll())
			if err != nil {
				raise(fmt.Sprintf("无法序列化配置: %v", err))
				return
			}

			responder(jsonData)

		case "refresh":
			if err := configManager.LoadJSON(); err != nil {
				raise(fmt.Sprintf("不能更新缓存: %v", err))
				return
			}

			responder([]byte("success"))

		default:
			raise(fmt.Sprintf("未知命令: %s", command))

		}
	})
}
