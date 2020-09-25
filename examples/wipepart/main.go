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
	fs := cmds.Filesystem{Device: diskPath}
	err := fs.Wipe()
	if err != nil {
		log.Fatalf("failed wiping disk: %v", err)
	}
}
