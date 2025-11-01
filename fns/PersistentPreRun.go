package fns

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"nodebus/cli"
	"nodebus/configm"

	ipcclient "github.com/runoneall/pgoipc/client"
	"github.com/spf13/cobra"
)

func PersistentPreRun(cmd *cobra.Command, args []string) {
	configManager := configm.GetManager()

	switch *cli.UseCfgCenter {

	case false:
		configManager.LoadJSON()

	case true:
		ipcclient.Connect("nodebus-cfgcenter", func(conn net.Conn) {
			if _, err := fmt.Fprintln(conn, "fetch"); err != nil {
				panic(fmt.Errorf("不能发送请求: %v", err))
			}

			resp, err := io.ReadAll(conn)
			if err != nil {
				panic(fmt.Errorf("不能接收响应: %v", err))
			}

			if err := configManager.LoadJSONFromReader(bytes.NewReader(resp)); err != nil {
				panic(fmt.Errorf("不能反序列化配置: %v", err))
			}
		})

	}
}
