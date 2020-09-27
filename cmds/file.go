// cerner_2^5_2020
package cmds

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type File struct {
	Path  string
	Owner string
	Group string
}

// Download a file
func (f *File) Download(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("HTTP Error %d", res.StatusCode)
	}
	file, err := os.Create(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
