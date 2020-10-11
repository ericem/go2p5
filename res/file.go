// cerner_2^5_2020
package res

import (
	"github.com/ericem/go2p5/node"
)

// Creates a Resource for managing the state of a file
func File(node *node.Node, url, path string, newState State) *Resource {
	curState := Absent
	file, found := node.GetFile(path)
	if found {
		curState = Present
	}
	fileRes := NewResource()
	if newState == Present && curState == Absent {
		fileRes.SetCommand(func() error {
			file.Path = path
			err := file.Download(url)
			return err
		})
	} else if newState == Absent && curState == Present {
		fileRes.SetCommand(func() error {
			err := file.Remove()
			return err
		})
	} else {
		fileRes.SetCommand(NoOp())
	}
	return &fileRes
}
