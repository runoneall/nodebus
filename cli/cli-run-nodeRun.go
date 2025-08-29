package cli

import (
	"context"
	"fmt"
	"nodebus/configm"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func nodeRun(cmd *cobra.Command, args []string) {
	manager := configm.GetManager()
	target_nodes := *SelectedNodes

	for _, name := range target_nodes {
		func() {
			fmt.Printf("\r\n--> 正在连接到 %s\r\n", name)

			fd := int(os.Stdin.Fd())
			oldState, err := term.MakeRaw(fd)
			if err != nil {
				fmt.Println("不能设置终端状态", err)
				return
			}
			defer term.Restore(fd, oldState)

			item, err := manager.ItemGet(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			client, err := ssh.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", item.SSH_Host, item.SSH_Port),
				&ssh.ClientConfig{
					User:            item.SSH_User,
					Auth:            []ssh.AuthMethod{ssh.Password(item.SSH_Password)},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
				},
			)

			if err != nil {
				fmt.Println("不能创建连接", err)
				return
			}
			defer client.Close()

			session, err := client.NewSession()
			if err != nil {
				fmt.Println("不能创建会话", err)
				return
			}
			defer session.Close()

			width, height, err := term.GetSize(fd)
			if err != nil {
				fmt.Println("不能获取终端大小", err)
				return
			}

			if err := session.RequestPty("xterm", height, width, ssh.TerminalModes{
				ssh.ECHO:          1,
				ssh.TTY_OP_ISPEED: 14400,
				ssh.TTY_OP_OSPEED: 14400,
			}); err != nil {
				fmt.Println("不能创建伪终端", err)
				return
			}

			winch_ctx, winch_cancel := context.WithCancel(context.Background())
			defer winch_cancel()

			go func() {
				ticker := time.NewTicker(100 * time.Millisecond)
				defer ticker.Stop()

				for {
					select {
					case <-ticker.C:
						width, height, err := term.GetSize(fd)
						if err != nil {
							continue
						}
						session.WindowChange(height, width)

					case <-winch_ctx.Done():
						return
					}
				}
			}()

			session.Stdin = os.Stdin
			session.Stdout = os.Stdout
			session.Stderr = os.Stderr

			cmd := strings.Join(args, " ")
			if cmd == "" {
				cmd = "sh"
			}

			if err := session.Run(cmd); err != nil {
				fmt.Println(err)
			}
		}()
	}

	fmt.Print("\r\n")
}
