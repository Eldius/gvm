package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	configDir = "~/.gvm"
)

func init() {
	_ = os.MkdirAll(GetVersionsDir(), os.ModePerm)
}

/*
Root returns the app config root folder
*/
func Root() string {
	cfgDir, err := homedir.Expand(configDir)
	if err != nil {
		log.Println("Failed to parse config folder")
		log.Panicln(err.Error())
	}
	return cfgDir
}

/*
GetWorkspaceDir returns the workspace dir
*/
func GetWorkspaceDir() string {
	return filepath.Join(Root(), "workspace")
}

/*
GetVersionsDir returns the versions dir
*/
func GetVersionsDir() string {
	return filepath.Join(GetWorkspaceDir(), "versions")
}

func GetVersionsPage() string {
	return viper.GetString("gvm.versions.page.url")
}
