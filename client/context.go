package client

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
	"sync"

	"github.com/feel-easy/hole/utils"
	"github.com/feel-easy/hole/utils/protocol"
	"github.com/gorilla/websocket"
)

type Context struct {
	sync.Mutex
	id    int
	name  string
	score int
	token string

	conn *protocol.Conn
}

type netConnector func(addr string) (*protocol.Conn, error)

var netConnectors = map[string]netConnector{
	"tcp": tcpConnect,
	"ws":  websocketConnect,
}

func NewContext(user LoginRespData) *Context {
	return &Context{
		id:    user.ID,
		name:  user.Name,
		score: user.Score,
		token: user.Token,
	}
}

func (c *Context) Connect(net string, addr string) error {
	if connector, ok := netConnectors[net]; ok {
		conn, err := connector(addr)
		if err != nil {
			return err
		}
		c.conn = conn
		return nil
	}
	return fmt.Errorf("unsupported net type: %s", net)
}

func (c *Context) Auth() error {
	return c.conn.Write(protocol.ObjectPacket(protocol.AuthInfo{
		ID:   c.id,
		Name: c.name,
	}))
}

func (c *Context) Listener() error {
	is := false
	utils.Async(func() {
		for {
			line, err := utils.ReadLine()
			if err != nil {
				log.Panic(err)
			}
			if !is {
				continue
			}
			c.print(fmt.Sprintf(cleanLine+"[%s@hole %s]# ", strings.TrimSpace(strings.ToLower(c.name)), "~"))
			err = c.conn.Write(protocol.Packet{
				Body: line,
			})
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	})
	return c.conn.Accept(func(packet protocol.Packet, conn *protocol.Conn) {
		data := string(packet.Body)
		switch data {
		case IsStart:
			if !is {
				c.print(fmt.Sprintf(cleanLine+"[%s@hole %s]# ", strings.TrimSpace(strings.ToLower(c.name)), "~"))
			}
			is = true
		case IsStop:
			if is {
				c.print(cleanLine)
			}
			is = false
		default:
			if is {
				c.print(cleanLine + data + fmt.Sprintf(cleanLine+"[%s@hole %s]# ", strings.TrimSpace(strings.ToLower(c.name)), "~"))
			} else {
				c.print(data)
			}
		}
	})
}

func (c *Context) print(str string) {
	c.Lock()
	defer c.Unlock()
	fmt.Print(str)
}

func tcpConnect(addr string) (*protocol.Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("tcp server error: %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("tcp server error: %v", err)
	}
	return protocol.Wrapper(protocol.NewTcpReadWriteCloser(conn)), nil
}

func websocketConnect(addr string) (*protocol.Conn, error) {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("ws server error: %v", err)
	}
	return protocol.Wrapper(protocol.NewWebsocketReadWriteCloser(conn)), nil
}
