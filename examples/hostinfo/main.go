package main

import (
	"fmt"

	"github.com/ericem/go2p5/cmds"
)

func main() {
	host := cmds.HostInfo()
	fmt.Printf("%8s: %-s\n", "Hostname", host.Name)
	fmt.Printf("%8s: %-s\n", "UUID", host.UUID)
	fmt.Printf("%8s: %-s\n", "Vendor", host.Vendor)
	fmt.Printf("%8s: %-s\n", "Model", host.Model)
	fmt.Printf("%8s: %-s\n", "Family", host.Family)
}
