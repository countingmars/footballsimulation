package data

import (
	"os"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)
var datapath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	datapath = filepath.Dir(b)
}
func LoadFile(filename string) []byte {
	path := absDataPath(filename)
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return bytes
}
func absDataPath(filename string) string {
	if 0 == strings.Index(filename, "/") {
		return datapath + filename
	} else {
		return datapath + "/" + filename
	}
}