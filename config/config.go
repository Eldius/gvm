package config

import (
	"log"

	"github.com/mitchellh/go-homedir"
)

/*
Root returns the app config root folder
*/
func Root() string {
	cfgDir, err := homedir.Expand("~/.gvm")
	if err != nil {
		log.Println("Failed to parse config folder")
		log.Panicln(err.Error())
	}
	return cfgDir
}
