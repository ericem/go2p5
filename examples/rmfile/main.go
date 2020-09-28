package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	f := cmds.File{Path: "examplefile.json"}
	err := f.Remove()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error removing file: %v\n", err)
		os.Exit(1)
	}
}
