package versions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

/*
GoVersion golang version representation
*/
type GoVersion struct {
	Name       string
	LinuxAmd64 string
	LinuxArm64 string
	Source     string
}

/*
Download downloads this version
*/
func (v *GoVersion) Download() (filePath string, err error) {
	fmt.Printf("Download package for version %s\n", v.Name)

	var resp *http.Response
	resp, err = http.Get(v.GetURL())
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	tmpFilePath, err := os.MkdirTemp("", fmt.Sprintf("%s.tar.gz", v.Name))
	if err != nil {
		return
	}

	// Create the file
	file, err := os.OpenFile(path.Join(tmpFilePath, fmt.Sprintf("%s.tar.gz", v.Name)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer func() {
		_ = file.Close()
	}()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	filePath = file.Name()
	return
}

func (v *GoVersion) GetURL() string {
	platform := fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
	switch {
	case strings.EqualFold(platform, linuxAmd64ArchName):
		return v.LinuxAmd64
	case strings.EqualFold(platform, linuxArm64ArchName):
		return v.LinuxArm64
	default:
		return v.LinuxAmd64
	}
}
