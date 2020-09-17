// cerner_2^5_2020
package cmds

import (
	"io/ioutil"
	"strings"
)

type OS struct {
	Name    string
	Version string
	ID      string
}

// OSinfo returns operating system version info
func OSInfo() (OS, error) {
	bs, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		return OS{}, err
	}
	lines := strings.Split(string(bs), "\n")
	fields := make(map[string]string)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "\"", "")
		if len(line) == 0 {
			continue
		}
		f := strings.Split(line, "=")
		fields[f[0]] = f[1]
	}
	return OS{Name: fields["NAME"], Version: fields["VERSION"], ID: fields["ID"]}, nil
}
