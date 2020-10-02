// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os/exec"
)

// Adds a new group to the system
func (g *Group) Add() error {
	_, err := exec.Command("groupadd", g.Name).Output()
	if err != nil {
		return fmt.Errorf("groupadd failed: %v", err)
	}
	return nil
}

// Removes a group from the system
func (g *Group) Remove() error {
	_, err := exec.Command("groupdel", g.Name).Output()
	if err != nil {
		return fmt.Errorf("groupdel failed: %v", err)
	}
	return nil
}
