package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Missing package name\n")
		os.Exit(1)
	}
	name := os.Args[1]
	pkg := cmds.Package{Name: name}
	err := pkg.Remove()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error removing package: %v\n", err)
		os.Exit(1)
	}
}
