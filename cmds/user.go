// cerner_2^5_2020
package cmds

import (
	"io/ioutil"
	"strings"
)

type User struct {
	Username string
	Home     string
	Shell    string
}

// Enumerates all users on a system
func Users() ([]User, error) {
	bs, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		return make([]User, 0), err
	}
	lines := strings.Split(string(bs), "\n")
	var users []User
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		f := strings.Split(line, ":")
		users = append(users, User{Username: f[0], Home: f[5], Shell: f[6]})
	}
	return users, nil
}
