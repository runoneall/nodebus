package fns

import (
	"context"
	"fmt"
	"net"
	"nodebus/cli"
	"nodebus/configm"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func NodeRun(args []string, isShell bool) {
	manager := configm.GetManager()
	target_nodes := *cli.SelectedNodes

	if *cli.IsAllNode {
		target_nodes = manager.ItemGetAllName()
	}

	for _, name := range target_nodes {
		func() {
			fmt.Printf("\r\n--> 正在连接到 %s\r\n", name)

			item, err := manager.ItemGet(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			client, err := ssh.Dial(
				"tcp",
				net.JoinHostPort(item.SSH_Host, item.SSH_Port),
				&ssh.ClientConfig{
					User: item.SSH_User,

					Auth: []ssh.AuthMethod{
						ssh.Password(item.SSH_Password),
						ssh.KeyboardInteractive(
							func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
								answers = make([]string, len(questions))
								for i := range questions {
									answers[i] = item.SSH_Password
								}
								return answers, nil
							},
						),
					},

					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Timeout:         3 * time.Minute,
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

			fd := int(os.Stdin.Fd())
			oldState, err := term.MakeRaw(fd)
			if err != nil {
				fmt.Println("不能设置终端状态", err)
				return
			}
			defer term.Restore(fd, oldState)

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

			switch isShell {

			case true:
				if err := session.Shell(); err != nil {
					fmt.Println(err)
				}

				if err := session.Wait(); err != nil {
					fmt.Println(err)
				}

			case false:
				cmd := strings.Join(args, " ")
				if cmd == "" {
					cmd = "sh"
				}

				if err := session.Run(cmd); err != nil {
					fmt.Println(err)
				}

			}
		}()
	}

	fmt.Print("\r\n")
}
