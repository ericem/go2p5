// cerner_2^5_2020
package cmds

import (
	"fmt"
	"io"
	"os"
)

type Sudoers struct {
	Name    string
	Content string
}

// Create a new sudoers file with arbitrary content
func (s *Sudoers) Add() error {
	sudoers := fmt.Sprintf("/etc/sudoers.d/%s", s.Name)
	f, _ := os.Create(sudoers)
	defer f.Close()
	_, err := io.WriteString(f, s.Content)
	if err != nil {
		return fmt.Errorf("failed writing sudoers: %v", err)
	}
	return nil
}

// Create a new sudoers file, granting username access without a password
func (s *Sudoers) NoPasswd(username string) error {
	s.Content = fmt.Sprintf("%%%s ALL=(ALL) NOPASSWD: ALL\n", username)
	err := s.Add()
	if err != nil {
		return fmt.Errorf("failed adding sudoers %s: %v", username, err)
	}
	return nil
}
