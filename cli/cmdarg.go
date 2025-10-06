package cli

var SelectedNodes *[]string
var IsAllNode *bool

var UseX11 *bool
var TrustX11 *bool

var UseCfgCenter *bool

var IsJSONOutput *bool
var SetJSONOutputIndent *int

func initCmdArg() {
	SelectedNodes = Cmd.PersistentFlags().StringSliceP("node", "n", []string{}, "指定要管理的节点")
	IsAllNode = Cmd.PersistentFlags().Bool("node-all", false, "指定管理全部节点")

	UseX11 = Cmd.PersistentFlags().Bool("x11", false, "启用 X11 转发")
	TrustX11 = Cmd.PersistentFlags().Bool("trust-x11", false, "完全信任 X11 (绕过 xauth)")

	UseCfgCenter = Cmd.PersistentFlags().Bool("cfgcenter", false, "从 cfgcenter 服务器拉取配置")

	AddCmd.Flags().String("name", "", "指定节点名称")
	AddCmd.Flags().String("host", "", "指定连接地址")
	AddCmd.Flags().String("port", "", "指定连接端口")
	AddCmd.Flags().String("user", "", "指定登录用户")
	AddCmd.Flags().String("pass", "", "指定登录密码")

	IsJSONOutput = ListCmd.Flags().BoolP("json", "j", false, "以 json 模式列出所有节点")
	SetJSONOutputIndent = ListCmd.Flags().IntP("indent", "i", 0, "设置 json 模式下的缩进")

	RunCmd.Flags().SetInterspersed(false)
	DockerCmd.Flags().SetInterspersed(false)
}
