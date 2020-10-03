package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	username := "test"
	sudoers := cmds.Sudoers{Name: username}
	err := sudoers.NoPasswd(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating sudoers: %v", err)
		os.Exit(1)
	}
}
