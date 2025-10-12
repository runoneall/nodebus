package main

import (
	"fmt"
	"nodebus/cli"
	"nodebus/fns"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := cli.Init()
	cmd.PersistentPreRun = fns.PersistentPreRun

	cli.AddCmd.Run = fns.NodeAdd
	cli.DelCmd.Run = fns.NodeDel
	cli.ListCmd.Run = fns.NodeList
	cli.RunCmd.Run = func(cmd *cobra.Command, args []string) { fns.NodeRun(args, false) }
	cli.DockerCmd.Run = func(cmd *cobra.Command, args []string) { fns.NodeRun(append([]string{"docker"}, args...), false) }
	cli.ShellCmd.Run = func(cmd *cobra.Command, args []string) { fns.NodeRun(args, true) }
	cli.CfgCenterCmd.Run = fns.CfgCenterServer
	cli.CfgShellCmd.Run = fns.CfgShell

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
