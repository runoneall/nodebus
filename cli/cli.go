package cli

import "github.com/spf13/cobra"

var SelectedNodes *[]string
var IsAllNode *bool

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

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "运行命令",
		Run:   nodeRun,
	}
	runCmd.Flags().SetInterspersed(false)

	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "操作 docker",
		Run: func(cmd *cobra.Command, args []string) {
			nodeRun(cmd, append([]string{"docker"}, args...))
		},
	}
	dockerCmd.Flags().SetInterspersed(false)

	cmd.AddCommand(
		addCmd,
		delCmd,
		listCmd,
		runCmd,
		dockerCmd,
	)

	return cmd
}
