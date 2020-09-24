package main

import (
	"log"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Missing disk path")
	}
	diskPath := os.Args[1]
	disk := cmds.Disk{Name: diskPath}
	err := disk.Partition(1, "8300")
	if err != nil {
		log.Fatalf("failed creating partition: %v", err)
	}
}
