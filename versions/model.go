package versions

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/*
GoVersion golang version representation
*/
type GoVersion struct {
	Name       string
	LinuxAmd64 string
	Source     string
}

func (v *GoVersion) Download() (filePath string, err error) {
	fmt.Printf("Download package for version %s\n", v.Name)

	var resp *http.Response
	resp, err = http.Get(v.LinuxAmd64)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	tmpFile, err := ioutil.TempFile("", fmt.Sprintf("%s.tar.gz", v.Name))
	if err != nil {
		return
	}
	// Create the file
	file, err := os.OpenFile(tmpFile.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	filePath = file.Name()
	return
}
