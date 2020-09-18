// cerner_2^5_2020
package parse

import (
	"strings"
)

// Parse a string containing key=value pairs into a map[string]string
func Pairs(s string) (map[string]string, error) {
	fields := make(map[string]string)
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return fields, nil
	}
	pairs := strings.Fields(s)
	for _, pair := range pairs {
		pair = strings.ReplaceAll(pair, "\"", "")
		p := strings.Split(pair, "=")
		key := strings.ToLower(p[0])
		value := p[1]
		fields[key] = value
	}

	return fields, nil
}
