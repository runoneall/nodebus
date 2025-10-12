package cli

import "github.com/spf13/cobra"

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加节点",
}

var DelCmd = &cobra.Command{
	Use:   "del",
	Short: "删除节点",
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有节点",
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "运行命令",
}

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "操作 docker",
}

var ShellCmd = &cobra.Command{
	Use:   "shell",
	Short: "登录远程 shell",
}

var CfgCenterCmd = &cobra.Command{
	Use:   "cfgcenter",
	Short: "集中式的管理节点配置",
}

var CfgShellCmd = &cobra.Command{
	Use:   "cfgshell",
	Short: "与 cfgcenter 交互的 shell",
}

func initSubCmd() {
	Cmd.AddCommand(
		AddCmd,
		DelCmd,
		ListCmd,
		RunCmd,
		DockerCmd,
		ShellCmd,
		CfgCenterCmd,
		CfgShellCmd,
	)
}
