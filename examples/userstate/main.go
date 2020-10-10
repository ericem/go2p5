package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ericem/go2p5/node"
	"github.com/ericem/go2p5/res"
)

const (
	keyurl  string = "https://raw.githubusercontent.com/hashicorp/vagrant/master/keys/vagrant.pub"
	keyname string = "id_rsa.pub"
)

func main() {
	var newState res.State
	present := flag.Bool("present", false, "ensure user is installed")
	absent := flag.Bool("absent", false, "ensure user is not installed")
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
	pkg := res.User(&node, "test", "testpass", keyurl, keyname, newState)
	err := pkg.Apply()
	if err != nil {
		fmt.Printf("error applying package state: %v\n", err)
		os.Exit(1)
	}
}
