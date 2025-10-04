package fns

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
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
	// get xauthority path
	xauthorityPath := os.Getenv("XAUTHORITY")
	if len(xauthorityPath) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			return errors.New("Xauthority not found: $XAUTHORITY, $HOME not set")
		}
		xauthorityPath = home + "/.Xauthority"
	}

	xa := xauth.XAuth{}
	xa.Display = os.Getenv("DISPLAY")

	cookie, err := xa.GetXAuthCookie(xauthorityPath, false)
	if err != nil {
		return err
	}

	// set x11-req Payload
	payload := x11Request{
		SingleConnection: false,
		AuthProtocol:     string("MIT-MAGIC-COOKIE-1"),
		AuthCookie:       string(cookie),
		ScreenNumber:     uint32(0),
	}

	// Send x11-req Request
	ok, err := session.SendRequest("x11-req", true, ssh.Marshal(payload))
	if err == nil && !ok {
		return errors.New("ssh: x11-req failed")
	} else {
		// Open HandleChannel x11
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

// x11Connect return net.Conn x11 socket.
func x11Connect(display string) (conn net.Conn, err error) {
	var conDisplay string

	protocol := "unix"

	if display[0] == '/' { // PATH type socket
		conDisplay = display
	} else if display[0] != ':' { // Forwarded display
		protocol = "tcp"
		if b, _, ok := strings.Cut(display, ":"); ok {
			conDisplay = fmt.Sprintf("%v:%v", b, getX11DisplayNumber(display)+6000)
		} else {
			conDisplay = display
		}
	} else { // /tmp/.X11-unix/X0
		conDisplay = fmt.Sprintf("/tmp/.X11-unix/X%v", getX11DisplayNumber(display))
	}

	return net.Dial(protocol, conDisplay)
}

// x11forwarder forwarding socket x11 data.
func x11forwarder(channel ssh.Channel) {
	conn, err := x11Connect(os.Getenv("DISPLAY"))

	if err != nil {
		return
	}

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
	conn.Close()
	channel.Close()
}

// getX11Display return X11 display number from env $DISPLAY
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
