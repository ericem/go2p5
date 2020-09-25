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
	fs := cmds.Filesystem{Device: diskPath, Type: "xfs"}
	err := fs.Format()
	if err != nil {
		log.Fatalf("failed formatting disk: %v", err)
	}
}
