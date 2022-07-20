package client

import (
	"fmt"
	"os"
	"time"

	"github.com/feel-easy/hole/utils"
	"github.com/feel-easy/hole/utils/logs"
)

type Client struct {
	ctx  *Context
	addr string
}

func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

func (s *Client) Start() error {
	fmt.Printf("Nickname: ")
	name, _ := utils.Readline()
	if len(os.Args) > 2 {
		s.addr = os.Args[2]
	}
	s.ctx = NewContext(LoginRespData{
		ID:       int(time.Now().UnixNano()),
		Name:     string(name),
		Score:    100,
		Username: string(name),
		Token:    "aeiou",
	})
	err := s.ctx.Connect("tcp", s.addr)
	if err != nil {
		logs.Error(err)
		return err
	}
	err = s.ctx.Auth()
	if err != nil {
		logs.Error(err)
		return err
	}
	return s.ctx.Listener()
}
