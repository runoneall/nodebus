package cli

import (
	"fmt"
	"nodebus/configm"

	"github.com/spf13/cobra"
)

func nodeDel(cmd *cobra.Command, args []string) {
	manager := configm.GetManager()
	target_deletes := *SelectedNodes

	if len(target_deletes) == 0 {
		fmt.Println("未选择任何节点")
		return
	}

	manager.ItemDel(
		target_deletes,
		func(name string) {
			fmt.Println("已删除", name)
		},
		func(name string) {
			fmt.Println("无法删除", name)
		},
	)

	if err := manager.SaveJSON(); err != nil {
		panic(err)
	}
}
