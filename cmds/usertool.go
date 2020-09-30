// cerner_2^5_2020
package cmds

import (
	"fmt"
	"io"
	"os/exec"
)

// Adds a user to the system
func (u *User) Add() error {
	_, err := exec.Command("useradd", "-m", "-s", u.Shell, u.Username).Output()
	if err != nil {
		return fmt.Errorf("useradd failed: %v", err)
	}
	return nil
}

func (u *User) SetPassword(pass string) error {
	cmd := exec.Command("passwd", "--stdin", u.Username)
	pipe, _ := cmd.StdinPipe()
	cmd.Start()
	io.WriteString(pipe, pass)
	cmdErr := cmd.Wait()
	if cmdErr != nil {
		return fmt.Errorf("passwd failed: %v", cmdErr)
	}
	return nil
}
