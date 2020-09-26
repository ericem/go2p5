// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os"
	"os/exec"
)

// Installs a yum package
func (p *Package) Install() error {
	out, err := exec.Command("yum", "install", "-y", "-q", p.Name).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", out)
		return fmt.Errorf("yum install failed: %v", err)
	}
	return nil
}

// Removes a yum package
func (p *Package) Remove() error {
	out, err := exec.Command("yum", "remove", "-y", "-q", p.Name).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", out)
		return fmt.Errorf("yum remove failed: %v", err)
	}
	return nil
}
