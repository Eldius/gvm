package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eldius/go-version-manager/config"
	"github.com/Eldius/go-version-manager/versions"
)

/*
Use sets up the version used
*/
func Use(version string) {
	for _, v := range versions.ListLocalVersions() {
		if versions.CompareVersions(version, v) {
			if !strings.HasPrefix(version, "go") {
				version = "go" + version
			}
			workspaceBinFolder := filepath.Join(config.GetWorkspaceDir(), "bin")
			//_ = os.MkdirAll(workspaceBinFolder, os.ModePerm)
			_ = os.Remove(workspaceBinFolder)
			versionBinDir := filepath.Join(config.GetVersionsDir(), version, "go", "bin")

			err := os.Symlink(versionBinDir, workspaceBinFolder)
			if err != nil {
				fmt.Printf("Failed to create symlink for the new active version")
				log.Panic(err)
			}
			updateRcFile(filepath.Join(config.GetHomeDir(), ".bashrc"))
			fmt.Println("version:", v)
			return
		}
	}
	fmt.Printf("Version not found locally...\nPlease install this version before run this command\n\ngo-version-manager install %s\n\n", version)
}

func updateRcFile(fileName string) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	s := string(b)
	pathUpdateStr := "export PATH=\"$HOME/.gvm/workspace/bin:$PATH\""

	if !strings.Contains(s, pathUpdateStr) {
		destFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}
		defer destFile.Close()
		if _, err := destFile.WriteString(fmt.Sprintf("\n# configure gvm path\n%s\n", pathUpdateStr)); err != nil {
			panic(err.Error())
		}
	}
}
