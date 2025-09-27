package cli

import (
	"github.com/spf13/cobra"
)

var SelectedNodes *[]string
var IsAllNode *bool

var UseCfgCenter *string
var CfgCenterAuth *string

var IsJSONOutput *bool
var SetJSONOutputIndent *int

var CfgCenterHost *string
var CfgCenterPort *string

func Init() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nodebus",
		Short: "在一处管理多台服务器",
	}

	SelectedNodes = cmd.PersistentFlags().StringSliceP("node", "n", []string{}, "指定要管理的节点")
	IsAllNode = cmd.PersistentFlags().Bool("node-all", false, "指定管理全部节点")

	UseCfgCenter = cmd.PersistentFlags().String("cfgcenter", "", "指定 cfgcenter 服务器")
	CfgCenterAuth = cmd.PersistentFlags().String("auth", "none", "连接到 cfgcenter 的认证字符串")

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

	cfgCenterCmd := &cobra.Command{
		Use:   "cfgcenter",
		Short: "集中式的管理节点配置",
		Run:   cfgCenterServer,
	}

	CfgCenterHost = cfgCenterCmd.Flags().String("host", "::", "指定 cfgcenter 的监听地址")
	CfgCenterPort = cfgCenterCmd.Flags().String("port", "32768", "指定 cfgcenter 的监听端口")

	cmd.AddCommand(
		addCmd,
		delCmd,
		listCmd,
		runCmd,
		dockerCmd,
		shellCmd,
		cfgCenterCmd,
	)

	return cmd
}
