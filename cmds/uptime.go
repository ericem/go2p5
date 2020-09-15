// cerner_2^5_2020
// Package cmds contains functions for common Linux systems administration
// commands.
package cmds

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Uptime returns the system uptime in Time.Duration type
func Uptime() (time.Duration, error) {
	out, err := ioutil.ReadFile("/proc/uptimes")
	if err != nil {
		return time.Duration(0), fmt.Errorf("reading /proc/uptime: %v", err)
	}

	sec, _ := strconv.ParseFloat(strings.Split(string(out), " ")[0], 64)
	if err != nil {
		return time.Duration(0), fmt.Errorf("parsing /proc/uptime: %v", err)
	}

	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", int64(sec)))
	if err != nil {
		return time.Duration(0), fmt.Errorf("parsing uptime duration: %v", err)
	}

	return duration, nil
}
