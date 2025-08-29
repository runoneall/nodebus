package main

import (
	"fmt"
	"nodebus/cli"
	"nodebus/configm"
	"os"
)

func main() {
	config_manager := configm.GetManager()
	config_manager.LoadJSON()

	if err := cli.Init().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
