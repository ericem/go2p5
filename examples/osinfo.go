package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	os, err := cmds.OSInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (%s) %s\n", os.Name, os.ID, os.Version)
}
