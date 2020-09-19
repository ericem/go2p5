package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	nics, err := cmds.Nics()
	if err != nil {
		log.Fatal(err)
	}
	for _, nic := range nics {
		fmt.Printf("%s %s/%d\n", nic.Name, nic.Addr, nic.Prefix)
	}
}
