package main

import (
	"fmt"
	"log"

	"github.com/ericem/go2p5/cmds"
)

func main() {

	users, err := cmds.Users()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Printf("%s\n", user.Username)
	}
}
