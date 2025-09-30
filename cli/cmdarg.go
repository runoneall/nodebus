package cli

var SelectedNodes *[]string
var IsAllNode *bool

var UseCfgCenter *string
var CfgCenterAuth *string

var IsJSONOutput *bool
var SetJSONOutputIndent *int

var CfgCenterHost *string
var CfgCenterPort *string

func initCmdArg() {
	SelectedNodes = Cmd.PersistentFlags().StringSliceP("node", "n", []string{}, "指定要管理的节点")
	IsAllNode = Cmd.PersistentFlags().Bool("node-all", false, "指定管理全部节点")

	UseCfgCenter = Cmd.PersistentFlags().String("cfgcenter", "", "指定 cfgcenter 服务器")
	CfgCenterAuth = Cmd.PersistentFlags().String("auth", "none", "认证字符串")

	AddCmd.Flags().String("name", "", "指定节点名称")
	AddCmd.Flags().String("host", "", "指定连接地址")
	AddCmd.Flags().String("port", "", "指定连接端口")
	AddCmd.Flags().String("user", "", "指定登录用户")
	AddCmd.Flags().String("pass", "", "指定登录密码")

	IsJSONOutput = ListCmd.Flags().BoolP("json", "j", false, "以 json 模式列出所有节点")
	SetJSONOutputIndent = ListCmd.Flags().IntP("indent", "i", 0, "设置 json 模式下的缩进")

	RunCmd.Flags().SetInterspersed(false)
	DockerCmd.Flags().SetInterspersed(false)

	CfgCenterHost = CfgCenterCmd.Flags().String("host", "::", "指定 cfgcenter 的监听地址")
	CfgCenterPort = CfgCenterCmd.Flags().String("port", "32768", "指定 cfgcenter 的监听端口")
}
