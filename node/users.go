// cerner_2^5_2020
package node

import (
	"os"

	"github.com/ericem/go2p5/cmds"
)

// Get the current User, empty user otherwise
func (n *Node) GetUser(username string) (*cmds.User, bool) {
	users, _ := cmds.Users()
	for i := range users {
		if users[i].Username == username {
			u := users[i]
			return &u, true
		}
	}
	return &cmds.User{}, false
}

// Check if a file or directory exists
func PathExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
