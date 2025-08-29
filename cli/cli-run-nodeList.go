package cli

import (
	"fmt"
	"nodebus/configm"
	"reflect"

	"github.com/spf13/cobra"
)

func nodeList(cmd *cobra.Command, args []string) {
	manager := configm.GetManager()
	all_node := manager.ItemGetAll()

	for name, item := range all_node {
		fmt.Printf("* %s\n", name)

		value := reflect.ValueOf(item)
		for i := 0; i < value.NumField(); i++ {
			k := value.Type().Field(i).Name
			v := value.Field(i).Interface()

			fmt.Printf("    - %s: %v\n", k, v)
		}
	}
}
