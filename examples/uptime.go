package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	uptime, err := cmds.Uptime()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current uptime is %v\n", uptime)
}
