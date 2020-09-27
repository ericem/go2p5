package main

import (
	"fmt"
	"os"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	url := "http://httpbin.org/json"
	f := cmds.File{Path: "examplefile.json"}
	err := f.Download(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Download complete")
}
