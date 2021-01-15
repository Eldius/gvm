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

const (
	configDirConst = "~/.gvm"
)

var (
	// Version app version
	Version string
	// BuildDate app build date
	BuildDate string
)

/*
Root returns the app config root folder
*/
func Root() string {
	configDir := viper.GetString("gvm.cfg.dir")
	if configDir == "" {
		configDir = configDirConst
	}
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

/*
GetHooksDir returns the hooks dir
*/
func GetHooksDir() string {
	return filepath.Join(GetWorkspaceDir(), "hooks")
}

/*
GetVersion returns the app version
*/
func GetVersion() string {
	return Version
}

/*
GetBuildDate returns the build date info
*/
func GetBuildDate() string {
	return BuildDate
}
