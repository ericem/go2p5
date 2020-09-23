// cerner_2^5_2020
package node

import (
	"encoding/json"

	"github.com/ericem/go2p5/cmds"
)

type Node struct {
	Host        cmds.Host         `json:"host"`
	OS          cmds.OS           `json:"os"`
	Disks       []cmds.Disk       `json:"disks"`
	Filesystems []cmds.Filesystem `json:"filesystems"`
	Nics        []cmds.Nic        `json:"nics"`
	Packages    []cmds.Package    `json:"packages"`
}

// Create a Node pre-populated with node information
func New() Node {
	var n = Node{}
	n.Host = cmds.HostInfo()
	n.OS, _ = cmds.OSInfo()
	n.Disks, _ = cmds.Disks()
	n.Filesystems, _ = cmds.Filesystems()
	n.Nics, _ = cmds.Nics()
	n.Packages, _ = cmds.RPMs()
	return n
}

// Convert a Node to JSON
func (n *Node) ToJson() []byte {
	j, _ := json.MarshalIndent(n, "", "  ")
	return j
}
