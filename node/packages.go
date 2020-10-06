// cerner_2^5_2020
package node

import (
	"github.com/ericem/go2p5/cmds"
)

// Check if a node has a package installed
func (n *Node) HasPackage(name string) bool {
	for i := range n.Packages {
		if n.Packages[i].Name == name {
			return true
		}
	}
	return false
}

// Check if a node has a specific version of a package installed
func (n *Node) HasPackageVersion(name string, version string) bool {
	for i := range n.Packages {
		if n.Packages[i].Name == name && n.Packages[i].Version == version {
			return true
		}
	}
	return false
}

// Get the current Package if installed, empty package otherwise
func (n *Node) GetPackage(name string) (*cmds.Package, bool) {
	for i := range n.Packages {
		if n.Packages[i].Name == name {
			p := n.Packages[i]
			return &p, true
		}
	}
	return &cmds.Package{}, false
}
