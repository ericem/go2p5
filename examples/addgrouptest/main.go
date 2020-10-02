package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	group := cmds.Group{Name: "test"}
	err := group.Add()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating group: %v", err)
		os.Exit(1)
	}
}
