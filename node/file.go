// cerner_2^5_2020
package node

import (
	"github.com/ericem/go2p5/cmds"
)

// Return a File if the path exists
func (n *Node) GetFile(path string) (*cmds.File, bool) {
	if PathExists(path) {
		return &cmds.File{Path: path}, true
	}
	return &cmds.File{}, false
}
