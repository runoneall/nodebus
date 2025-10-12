package fns

import (
	"bufio"
	"fmt"
	"nodebus/cli"
	"nodebus/ipc"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func CfgShell(cmd *cobra.Command, args []string) {
	cmds := *cli.CfgShellExec

	conn := ipc.Connect("cfgcenter")
	defer conn.Close()

	execute := func(input string) string {
		if err := conn.Send([]byte(input)); err != nil {
			panic(fmt.Errorf("不能发送请求: %v", err))
		}

		output, err := conn.Recv()
		if err != nil {
			panic(fmt.Errorf("不能获取输出: %v", err))
		}

		return string(output)
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

			fmt.Println(execute(input))
		}
	}

	switch len(cmds) {

	case 0:
		startShell()

	default:
		for _, command := range cmds {
			fmt.Println(execute(command))
		}

	}
}
