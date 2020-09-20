package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {

	filesystems, err := cmds.Filesystems()
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range filesystems {
		fmt.Printf("%8s: %-s\n", "Device", fs.Device)
		fmt.Printf("%8s: %-s\n", "Mount", fs.Mount)
		fmt.Printf("%8s: %s\n", "Type", fs.Type)
		fmt.Printf("%8s: %s\n", "Options", fs.Options)
		fmt.Println("")
	}

}
