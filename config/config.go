package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func init() {
	_ = os.MkdirAll(GetVersionsDir(), os.ModePerm)
}

/*
Root returns the app config root folder
*/
func Root() string {
	configDir := viper.GetString("gvm.cfg.dir")
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

/*
GetVersionsPage returns the versions page
*/
func GetVersionsPage() string {
	return viper.GetString("gvm.versions.page.url")
}

/*
GetHomeDir returns the home dir
*/
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Failed to find user home dir")
		log.Println("Failed to find user home dir")
		log.Panic(err.Error())
	}
	return home
}
