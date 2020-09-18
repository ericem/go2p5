package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/ericem/go2p5/parse"
)

type Filesystem struct {
	Device  string
	Mount   string
	Type    string
	Options string
	Label   string
	UUID    string
}

func Filesystems(fstype string) ([]Filesystem, error) {
	out, err := exec.Command("findmnt", "-P", "-t", fstype, "-o",
		"source,target,fstype,options,label,uuid").Output()
	if err != nil {
		return make([]Filesystem, 0), fmt.Errorf("failed lsblk command: %v", err)
	}
	var fs []Filesystem
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) == 0 {
			continue
		}
		f, err := parse.Pairs(line)
		if err != nil {
			return make([]Filesystem, 0), fmt.Errorf("failed parsing: %v", err)
		}
		if len(f) == 0 {
			continue
		}
		fs = append(fs, Filesystem{Device: f["source"], Mount: f["target"],
			Type: f["fstype"], Options: f["options"], Label: f["label"],
			UUID: f["uuid"]})
	}
	return fs, nil
}

func XFS() ([]Filesystem, error) {
	return Filesystems("xfs")
}

func main() {

	filesystems, err := XFS()
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range filesystems {
		fmt.Printf("%s is mounted on %s as %s\n", fs.Device, fs.Mount, fs.Type)
	}

}
