package cli

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "nodebus",
	Short: "在一处管理多台服务器",
}

func Init() *cobra.Command {
	initSubCmd()
	initCmdArg()

	return Cmd
}
