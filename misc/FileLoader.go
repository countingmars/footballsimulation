package misc

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
)

func LoadFile(filename string) []byte {
	path, _ := filepath.Abs(filename)
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return bytes
}
