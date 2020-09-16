// cerner_2^5_2020
package cmds

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Disk struct {
	Name string
	Size int64
	Type string
}

// Uptime enumerates all disks on a system
func Disks() ([]Disk, error) {
	cmd := exec.Command("lsblk", "-bdnp", "-o", "name,size,tran")
	pipe, _ := cmd.StdoutPipe()
	cmd.Start()
	input := bufio.NewScanner(pipe)
	var disks []Disk
	for input.Scan() {
		fields := strings.Fields(input.Text())
		size, _ := strconv.ParseInt(fields[1], 0, 64)
		disk := Disk{Name: fields[0], Size: size, Type: fields[2]}
		disks = append(disks, disk)
	}
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		return make([]Disk, 0), fmt.Errorf("waiting lsblk: %v", cmdErr)
	}
	return disks, nil
}
