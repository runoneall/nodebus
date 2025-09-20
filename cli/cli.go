package cli

import "github.com/spf13/cobra"

var SelectedNodes *[]string
var IsAllNode *bool

var IsJSONOutput *bool
var SetJSONOutputIndent *int

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodebus",
		Short: "在一处管理多台服务器",
	}

	SelectedNodes = cmd.PersistentFlags().StringSliceP("node", "n", []string{}, "指定要管理的节点")
	IsAllNode = cmd.PersistentFlags().Bool("node-all", false, "指定管理全部节点")

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "添加节点",
		Run:   nodeAdd,
	}

	addCmd.Flags().String("name", "", "指定节点名称")
	addCmd.Flags().String("host", "", "指定连接地址")
	addCmd.Flags().String("port", "", "指定连接端口")
	addCmd.Flags().String("user", "", "指定登录用户")
	addCmd.Flags().String("pass", "", "指定登录密码")

	delCmd := &cobra.Command{
		Use:   "del",
		Short: "删除节点",
		Run:   nodeDel,
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "列出所有节点",
		Run:   nodeList,
	}

	IsJSONOutput = listCmd.Flags().BoolP("json", "j", false, "以 json 模式列出所有节点")
	SetJSONOutputIndent = listCmd.Flags().IntP("indent", "i", 0, "设置 json 模式下的缩进")

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "运行命令",
		Run: func(cmd *cobra.Command, args []string) {
			nodeRun(args, false)
		},
	}
	runCmd.Flags().SetInterspersed(false)

	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "操作 docker",
		Run: func(cmd *cobra.Command, args []string) {
			nodeRun(append([]string{"docker"}, args...), false)
		},
	}
	dockerCmd.Flags().SetInterspersed(false)

	shellCmd := &cobra.Command{
		Use:   "shell",
		Short: "登录远程 shell",
		Run: func(cmd *cobra.Command, args []string) {
			nodeRun(args, true)
		},
	}

	cmd.AddCommand(
		addCmd,
		delCmd,
		listCmd,
		runCmd,
		dockerCmd,
		shellCmd,
	)

	return cmd
}
