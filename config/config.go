package config

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const (
	configDir = "~/.gvm"
)

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

func GetWorkspaceDir() string {
	return filepath.Join(Root(), "workspace")
}

func GetVersionsDir() string {
	return filepath.Join(GetWorkspaceDir(), "versions")
}
