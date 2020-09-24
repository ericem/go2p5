// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os/exec"
)

// Wipes the partition table on a disk block device
func (d *Disk) Wipe() error {
	_, err := exec.Command("sgdisk", "-Z", d.Name).Output()
	if err != nil {
		return fmt.Errorf("sgdisk failed: %v", err)
	}
	return nil
}

// Creates a partition on a disk block device
func (d *Disk) Partition(num int, partType string) error {
	_, err := exec.Command("sgdisk", "-g", "-n", fmt.Sprintf("%d:0:0", num),
		"-t", fmt.Sprintf("%d:%s", num, partType), d.Name).Output()
	if err != nil {
		return fmt.Errorf("sgdisk failed: %v", err)
	}
	return nil
}
