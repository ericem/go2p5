// cerner_2^5_2020
package cmds

import (
	"io/ioutil"
	"os"
	"strings"
)

type Host struct {
	Name   string
	UUID   string
	Model  string
	Family string
	Vendor string
}

// HostInfo returns basic information about the host
func HostInfo() Host {
	var h = Host{}
	h.Name, _ = os.Hostname()
	uuid, _ := ioutil.ReadFile("/sys/class/dmi/id/product_uuid")
	h.UUID = strings.TrimSpace(string(uuid))
	model, _ := ioutil.ReadFile("/sys/class/dmi/id/product_name")
	h.Model = strings.TrimSpace(string(model))
	family, _ := ioutil.ReadFile("/sys/class/dmi/id/product_family")
	h.Family = strings.TrimSpace(string(family))
	vendor, _ := ioutil.ReadFile("/sys/class/dmi/id/board_vendor")
	h.Vendor = strings.TrimSpace(string(vendor))
	return h
}
