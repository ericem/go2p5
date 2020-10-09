// cerner_2^5_2020
package res

import (
	"github.com/ericem/go2p5/node"
)

// Creates a Resource for managing the state of a package
func User(node *node.Node, username, pass, keyurl, keyname string, newState State) *Resource {
	curState := Absent
	user, found := node.GetUser(username)
	if found {
		curState = Present
	}
	pkgRes := NewResource()
	if newState == Present && curState == Absent {
		pkgRes.SetCommand(func() error {
			user.Username = username
			err := user.Add()
			err = user.SetPassword(pass)
			err = user.SetPubKey(keyurl, keyname)
			return err
		})
	} else if newState == Absent && curState == Present {
		pkgRes.SetCommand(func() error {
			err := user.Remove()
			return err
		})
	} else {
		pkgRes.SetCommand(NoOp())
	}
	return &pkgRes
}
