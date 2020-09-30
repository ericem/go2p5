package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	user := cmds.User{Username: "test", Home: "/home/test", Shell: "/bin/bash"}
	err := user.Add()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating user: %v", err)
		os.Exit(1)
	}
	user.SetPassword("tooeasy")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error setting password: %v", err)
		os.Exit(1)
	}
}
