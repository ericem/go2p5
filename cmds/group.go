// cerner_2^5_2020
package cmds

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Group struct {
	Name string
	ID   int
}

// Enumerates all users on a system
func Groups() ([]Group, error) {
	bs, err := ioutil.ReadFile("/etc/group")
	if err != nil {
		return make([]Group, 0), err
	}
	lines := strings.Split(string(bs), "\n")
	var users []Group
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		f := strings.Split(line, ":")
		id, _ := strconv.Atoi(f[2])
		users = append(users, Group{Name: f[0], ID: id})
	}
	return users, nil
}
