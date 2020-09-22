package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	rpms, err := cmds.RPMs()
	if err != nil {
		log.Fatalf("error reading rpms: %v", err)
	}
	for _, pkg := range rpms {
		fmt.Printf("%s %s\n", pkg.Name, pkg.Version)
	}
}
