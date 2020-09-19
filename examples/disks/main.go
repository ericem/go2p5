package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	disks, err := cmds.Disks()
	if err != nil {
		log.Fatal(err)
	}
	for _, disk := range disks {
		fmt.Printf("Disk %s has %d bytes\n", disk.Name, disk.Size)
	}
}
