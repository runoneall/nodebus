package cli

import (
	"bufio"
	"fmt"
	"net"
	"nodebus/configm"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func nodeAdd(cmd *cobra.Command, args []string) {
	manager := configm.GetManager()

	require_input := func(prompt string, verifys []func(input string) bool) string {
		for {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print(prompt, ": ")
			input, err := reader.ReadString('\n')

			if err != nil {
				panic(err)
			}
			input = strings.TrimSpace(input)

			verified_cont := 0
			for _, verify := range verifys {
				if !verify(input) {
					break
				}
				verified_cont += 1
			}

			if verified_cont == len(verifys) {
				return input
			}
		}
	}

	general_verify := func(input string) bool {

		if input == "" {
			fmt.Println("不能为空")
			return false
		}

		if strings.Contains(input, " ") {
			fmt.Println("不能包含空格")
			return false
		}

		return true
	}

	type verifys []func(input string) bool

	item_name := require_input(
		"节点名称 (node0)",
		verifys{
			general_verify,
			func(input string) bool {
				if manager.ItemExists(input) {
					fmt.Println("节点名称已存在")
					return false
				}
				return true
			},
		},
	)

	item_ssh_host := require_input(
		"SSH 主机地址 (127.0.0.1)",
		verifys{
			general_verify,
			func(input string) bool {
				is_ip := net.ParseIP(input)
				addrs, _ := net.LookupHost(input)

				if is_ip == nil && len(addrs) == 0 {
					fmt.Println("请输入正确的 IP 地址或域名")
					return false
				}
				return true
			},
		},
	)

	_item_ssh_port := require_input(
		"SSH 连接端口 (22)",
		verifys{
			general_verify,
			func(input string) bool {
				if _, err := strconv.Atoi(input); err != nil {
					fmt.Println("必须是数字")
					return false
				}
				return true
			},
		},
	)
	item_ssh_port, _ := strconv.Atoi(_item_ssh_port)

	item_ssh_user := require_input(
		"SSH 登录用户 (root)",
		verifys{
			general_verify,
		},
	)

	item_ssh_password := require_input(
		"SSH 密码 (abcde12345)",
		verifys{
			func(input string) bool {
				if input == "" {
					fmt.Println("将使用无密码登录")
				}
				return true
			},
		},
	)

	if strings.Contains(item_ssh_host, ":") {
		item_ssh_host = "[" + item_ssh_host + "]"
	}

	manager.ItemAdd(configm.Item{
		Name:         item_name,
		SSH_Host:     item_ssh_host,
		SSH_Port:     item_ssh_port,
		SSH_User:     item_ssh_user,
		SSH_Password: item_ssh_password,
	})

	if err := manager.SaveJSON(); err != nil {
		panic(err)
	}

	fmt.Println("节点添加成功")
}
