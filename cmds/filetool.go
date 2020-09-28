// cerner_2^5_2020
package cmds

import (
	"fmt"
	"os"
)

// Remove a file
func (f *File) Remove() error {
	err := os.Remove(f.Path)
	if err != nil {
		return fmt.Errorf("error removing file: %v\n", err)
	}
	return nil
}
