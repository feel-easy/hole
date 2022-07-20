/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/feel-easy/hole/client"
	"github.com/feel-easy/hole/utils/logs"
)

var (
	address string
)

const HELP = ""

func init() {
	flag.StringVar(&address, "addr", "127.0.0.1:9999", "服务器的地址")
}

func main() {
	flag.Parse()
	if len(os.Args) > 0 {
		for _, arg := range os.Args {
			if arg == "-help" {
				fmt.Print(HELP)
				os.Exit(0)
			}
		}
	}
	logs.Error(client.NewClient(address).Start())
}
