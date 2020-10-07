// cerner_2^5_2020
package res

import (
	"github.com/ericem/go2p5/node"
)

// Creates a Resource for managing the state of a package
func Package(node *node.Node, name string, newState State) *Resource {
	curState := Absent
	pkg, found := node.GetPackage(name)
	if found {
		curState = Present
	}
	pkgRes := NewResource()
	if newState == Present && curState == Absent {
		pkgRes.SetCommand(func() error {
			pkg.Name = name
			err := pkg.Install()
			return err
		})
	} else if newState == Absent && curState == Present {
		pkgRes.SetCommand(func() error {
			err := pkg.Remove()
			return err
		})
	} else {
		pkgRes.SetCommand(NoOp())
	}
	return &pkgRes
}
