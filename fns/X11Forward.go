package fns

import (
	"fmt"
	"io"
	"net"
	"nodebus/cli"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	xauth "github.com/blacknon/go-x11auth"
	"golang.org/x/crypto/ssh"
)

type x11Request struct {
	SingleConnection bool
	AuthProtocol     string
	AuthCookie       string
	ScreenNumber     uint32
}

func X11Forward(
	client *ssh.Client,
	session *ssh.Session,
) error {
	xauthorityPath := os.Getenv("XAUTHORITY")
	if xauthorityPath == "" {
		home := os.Getenv("HOME")
		if home == "" {
			return fmt.Errorf(".Xauthority not found: $XAUTHORITY, $HOME not set")
		}
		xauthorityPath = filepath.Join(home, ".Xauthority")
	}

	xa := xauth.XAuth{}
	xa.Display = os.Getenv("DISPLAY")

	cookie, err := xa.GetXAuthCookie(xauthorityPath, *cli.TrustX11)
	if err != nil {
		return err
	}

	payload := x11Request{
		SingleConnection: false,
		AuthProtocol:     "MIT-MAGIC-COOKIE-1",
		AuthCookie:       cookie,
		ScreenNumber:     0,
	}

	ok, err := session.SendRequest("x11-req", true, ssh.Marshal(payload))
	if err == nil && !ok {
		return fmt.Errorf("不能发送 X11 请求: %v", err)

	} else {
		x11channels := client.HandleChannelOpen("x11")

		go func() {
			for ch := range x11channels {
				channel, _, err := ch.Accept()
				if err != nil {
					continue
				}

				go x11forwarder(channel)
			}
		}()
	}

	return err
}

func x11Connect(display string) (net.Conn, error) {
	var conDisplay string

	protocol := "unix"

	if display[0] == '/' {
		conDisplay = display

	} else if display[0] != ':' {
		protocol = "tcp"

		if b, _, ok := strings.Cut(display, ":"); ok {
			conDisplay = fmt.Sprintf("%v:%v", b, getX11DisplayNumber(display)+6000)

		} else {
			conDisplay = display
		}

	} else {
		conDisplay = fmt.Sprintf("/tmp/.X11-unix/X%v", getX11DisplayNumber(display))
	}

	return net.Dial(protocol, conDisplay)
}

func x11forwarder(channel ssh.Channel) {
	defer channel.Close()

	conn, err := x11Connect(os.Getenv("DISPLAY"))
	if err != nil {
		return
	}
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		io.Copy(conn, channel)
		conn.Close()
		wg.Done()
	}()
	go func() {
		io.Copy(channel, conn)
		channel.CloseWrite()
		wg.Done()
	}()

	wg.Wait()
}

func getX11DisplayNumber(display string) int {
	colonIdx := strings.LastIndex(display, ":")
	dotIdx := strings.LastIndex(display, ".")

	if colonIdx < 0 {
		return 0
	}

	if dotIdx < 0 {
		dotIdx = len(display)
	}

	i, err := strconv.Atoi(display[colonIdx+1 : dotIdx])
	if err != nil {
		return 0
	}

	return i
}
