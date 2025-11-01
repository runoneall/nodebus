package fns

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"nodebus/cli"
	"os"
	"strings"

	ipcclient "github.com/runoneall/pgoipc/client"
	"github.com/spf13/cobra"
)

func CfgShell(cmd *cobra.Command, args []string) {
	cmds := *cli.CfgShellExec

	execute := func(input string) {
		ipcclient.Connect("nodebus-cfgcenter", func(conn net.Conn) {
			if _, err := fmt.Fprintln(conn, input); err != nil {
				panic(fmt.Errorf("不能发送请求: %v", err))
			}

			io.Copy(os.Stdout, conn)
			fmt.Println("")
		})
	}

	startShell := func() {
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Print("cfgshell> ")

			input, err := reader.ReadString('\n')
			if err != nil {
				panic(fmt.Errorf("不能读取输入: %v", err))
			}
			input = strings.TrimSpace(input)

			if input == "" {
				continue
			}

			execute(input)
		}
	}

	switch len(cmds) {

	case 0:
		startShell()

	default:
		for _, command := range cmds {
			execute(command)
		}

	}

}
