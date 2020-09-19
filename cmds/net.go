// cerner_2^5_2020
package cmds

import (
	"encoding/json"
	"os/exec"
)

type Nic struct {
	Name   string
	Addr   string
	Prefix int
}

// Nics enumerates all network interfaces on a system
func Nics() ([]Nic, error) {
	out, _ := exec.Command("ip", "-4", "-j", "addr", "show").Output()
	var nics []Nic
	var ipdata []interface{}
	_ = json.Unmarshal(out, &ipdata)
	for _, n := range ipdata {
		switch n := n.(type) {
		case map[string]interface{}:
			addrs, _ := n["addr_info"].([]interface{})
			a, _ := addrs[0].(map[string]interface{})
			name, _ := n["ifname"].(string)
			addr, _ := a["local"].(string)
			prfx, _ := a["prefixlen"].(float64)
			nics = append(nics, Nic{Name: name, Addr: addr, Prefix: int(prfx)})
		}
	}
	return nics, nil
}
