// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os/exec"
	"strings"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
	Arch    string `json:"arch"`
}

// Queries the rpm database for all installed packages
func RPMs() ([]Package, error) {
	out, err := exec.Command("rpm", "-qa", "--qf", `%{NAME} %{VERSION} %{RELEASE} %{ARCH}\n`).Output()
	if err != nil {
		return make([]Package, 0), fmt.Errorf("failed rpm command: %v", err)
	}
	var rpms []Package
	for _, line := range strings.Split(string(out), "\n") {
		f := strings.Fields(strings.TrimSpace(line))
		if len(f) != 4 {
			continue
		}
		rpms = append(rpms, Package{Name: string(f[0]), Version: string(f[1]),
			Release: string(f[2]), Arch: string(f[3])})
	}
	return rpms, nil
}
