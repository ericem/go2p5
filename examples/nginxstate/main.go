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
	present := flag.Bool("present", false, "ensure nginx is installed")
	absent := flag.Bool("absent", false, "ensure nginx is not installed")
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
	pkg := res.Package(&node, "nginx", newState)
	err := pkg.Apply()
	if err != nil {
		fmt.Printf("error applying package state: %v\n", err)
		os.Exit(1)
	}
}
