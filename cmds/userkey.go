// cerner_2^5_2020
package cmds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// Remove a user
func (u *User) Remove() error {
	_, err := exec.Command("userdel", "-r", u.Username).Output()
	if err != nil {
		return fmt.Errorf("userdel failed: %v", err)
	}
	return nil
}

// Download a public key for a user
func (u *User) SetPubKey(url string, name string) error {
	res, _ := http.Get(url)
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP Error %d, couldn't download key", res.StatusCode)
	}
	defer res.Body.Close()
	keydir := fmt.Sprintf("/home/%s/.ssh", u.Username)
	err := os.Mkdir(keydir, 0600)
	path := fmt.Sprintf("%s/%s", keydir, name)
	f, err := os.Create(path)
	defer f.Close()
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return fmt.Errorf("couldn't create key file: %v", err)
	}
	return nil
}
