// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/ericem/go2p5/parse"
)

const fstypes = "ext4,xfs"

type Filesystem struct {
	Device  string
	Mount   string
	Type    string
	Options string
}

func Filesystems() ([]Filesystem, error) {
	out, err := exec.Command("findmnt", "-P", "-t", fstypes, "-o", "source,target,fstype,options").Output()
	if err != nil {
		return make([]Filesystem, 0), fmt.Errorf("failed lsblk command: %v", err)
	}
	var fs []Filesystem
	for _, line := range strings.Split(string(out), "\n") {
		f, err := parse.Pairs(line)
		if err != nil {
			return make([]Filesystem, 0), fmt.Errorf("failed parsing: %v", err)
		}
		if len(f) > 0 {
			fs = append(fs, Filesystem{Device: f["source"], Mount: f["target"], Type: f["fstype"], Options: f["options"]})
		}
	}
	return fs, nil
}
