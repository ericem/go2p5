package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {

	groups, err := cmds.Groups()
	if err != nil {
		log.Fatal(err)
	}
	for _, group := range groups {
		fmt.Printf("%s\n", group.Name)
	}
}
