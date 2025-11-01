package fns

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"nodebus/configm"
	"strings"
	"time"

	ipcserver "github.com/runoneall/pgoipc/server"
	"github.com/spf13/cobra"
)

func CfgCenterServer(cmd *cobra.Command, args []string) {
	configManager := configm.GetManager()

	ipcserver.Serv("nodebus-cfgcenter", func(conn net.Conn) {
		reader := bufio.NewReader(conn)

		response := func(data []byte) {
			conn.Write(data)
		}

		command, err := reader.ReadString('\n')
		if err != nil {
			response(fmt.Appendf([]byte{}, "不能读取请求: %v", err))
			return
		}
		command = strings.TrimSpace(command)

		switch command {

		case "fetch":
			jsonData, err := json.Marshal(configManager.ItemGetAll())
			if err != nil {
				response(fmt.Appendf([]byte{}, "无法序列化配置: %v", err))
				return
			}

			response(jsonData)
			return

		case "refresh":
			if err := configManager.LoadJSON(); err != nil {
				response(fmt.Appendf([]byte{}, "不能更新缓存: %v", err))
				return
			}

			response([]byte("success"))
			return

		case "steamtest":
			for i := range 10 {
				response(fmt.Appendf([]byte{}, "Message %d (and wait 1s)\n", i))
				time.Sleep(1 * time.Second)
			}

		default:
			response(fmt.Appendf([]byte{}, "未知命令: %s\n", command))
			response([]byte("可用命令: fetch | refresh | steamtest"))
			return

		}
	})
}
