package main

import (
	"fmt"

	"github.com/ericem/go2p5/node"
)

func main() {
	n := node.New()
	fmt.Printf("%s\n", n.ToJson())
}
