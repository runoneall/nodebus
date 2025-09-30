package fns

import (
	"encoding/json"
	"fmt"
	"nodebus/cli"
	"nodebus/configm"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

func NodeList(cmd *cobra.Command, args []string) {
	manager := configm.GetManager()
	all_node := manager.ItemGetAll()

	switch *cli.IsJSONOutput {

	case true:
		indent := *cli.SetJSONOutputIndent

		var data []byte
		var jsonErr error

		switch indent {

		case 0:
			data, jsonErr = json.Marshal(all_node)

		default:
			data, jsonErr = json.MarshalIndent(all_node, "", strings.Repeat(" ", indent))

		}

		if jsonErr != nil {
			panic(fmt.Errorf("不能序列化 json: %v", jsonErr))
		}

		fmt.Println(string(data))

	case false:
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
}
