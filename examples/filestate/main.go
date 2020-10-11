package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ericem/go2p5/node"
	"github.com/ericem/go2p5/res"
)

func main() {
	var newState res.State
	present := flag.Bool("present", false, "ensure file is present")
	absent := flag.Bool("absent", false, "ensure file is absent")
	flag.Parse()
	if *present == true && *absent == false {
		newState = res.Present
	} else if *present == false && *absent == true {
		newState = res.Absent
	} else {
		flag.Usage()
		os.Exit(1)
	}
	node := node.New()
	file := res.File(&node, "http://httpbin.org/json", "/etc/example.json", newState)
	err := file.Apply()
	if err != nil {
		fmt.Printf("error applying file state: %v\n", err)
		os.Exit(1)
	}
}
