// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os/exec"
)

// Format a disk partition with a filesystem
func (f *Filesystem) Format() error {
	out, err := exec.Command("mkfs", "-t", f.Type, f.Device).CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		return fmt.Errorf("mkfs failed: %v", err)
	}
	return nil
}

// Wipe a filesystem from a disk partition
func (f *Filesystem) Wipe() error {
	out, err := exec.Command("wipefs", "-a", f.Device).CombinedOutput()
	if err != nil {
		fmt.Printf("%s\n", out)
		return fmt.Errorf("wipe failed: %v", err)
	}
	return nil
}
