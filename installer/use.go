package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eldius/gvm/config"
	"github.com/Eldius/gvm/versions"
)

/*
Use sets up the version used
*/
func Use(version string) error {
	for _, v := range versions.ListLocalVersions() {
		if versions.CompareVersions(version, v) {
			if !strings.HasPrefix(version, "go") {
				version = "go" + version
			}
			workspaceBinFolder := filepath.Join(config.GetWorkspaceDir(), "bin")
			workspacePathFolder := filepath.Join(config.GetWorkspaceDir(), "path")
			//_ = os.MkdirAll(workspaceBinFolder, os.ModePerm)
			_ = os.Remove(workspaceBinFolder)
			_ = os.Remove(workspacePathFolder)
			versionBinDir := filepath.Join(config.GetVersionsDir(), version, "go", "bin")
			pathDir := filepath.Join(config.GetVersionsDir(), version, "path")
			_ = os.MkdirAll(pathDir, os.ModePerm)

			err := os.Symlink(versionBinDir, workspaceBinFolder)
			if err != nil {
				fmt.Printf("Failed to create symlink for the new active version")
				log.Panic(err)
				return err
			}
			err = os.Symlink(pathDir, workspacePathFolder)
			if err != nil {
				fmt.Printf("Failed to create symlink for the new active version")
				log.Panic(err)
				return err
			}
			home := config.GetHomeDir()
			for _, f := range []string{filepath.Join(home, ".bashrc"), filepath.Join(home, ".zshrc")} {
				_, err := os.Stat(f)
				if err == nil {
					updateRcFile(f)
				}
			}
			fmt.Println("version:", v)
			return nil
		}
	}
	//fmt.Printf("Version not found locally...\nPlease install this version before run this command\n\ngo-version-manager install %s\n\n", version)
	return fmt.Errorf("Version not found locally...\nPlease install this version before run this command\n\ngo-version-manager install %s\n\n", version)
}

func updateRcFile(fileName string) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	s := string(b)
	pathUpdateStr := "export PATH=\"$HOME/.gvm/workspace/bin:$HOME/.gvm/workspace/path/bin:$PATH\""
	goPathStr := "export GOPATH=\"$HOME/.gvm/workspace/path\""

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

	if !strings.Contains(s, goPathStr) {
		destFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}
		defer destFile.Close()
		if _, err := destFile.WriteString(fmt.Sprintf("\n# configure go path\n%s\n", goPathStr)); err != nil {
			panic(err.Error())
		}
	}
}
