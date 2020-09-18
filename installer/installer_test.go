package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/Eldius/go-version-manager/versions"
	"github.com/spf13/viper"
)

var tmpDir string

func init() {
	var err error
	tmpDir, err = ioutil.TempDir("", "install")
	if err != nil {
		log.Panic("Error creating temp dir\n", err.Error())
	}
	viper.SetDefault("gvm.cfg.dir", tmpDir)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("falha carregando arquivo de configuracoes\n", err.Error())
	}
}
func TestUnpack(t *testing.T)  {
	
	v := &versions.GoVersion{
		Name: "1.2.3",
	}
	log.Printf("testing...")
	unpack("test_data/test.tar.gz", v)
	wsDir := filepath.Join(tmpDir, "workspace")
	if fs, err := os.Stat(wsDir); err != nil {
		t.Errorf("Could not stat workspace folder at '%s'...", wsDir)
	} else if !fs.IsDir() {
		t.Errorf("'%s' must be a folder...", wsDir)
	}

	versionDir := filepath.Join(wsDir, "versions", v.Name)
	if fs, err := os.Stat(versionDir); err != nil {
		t.Errorf("Could not stat version folder at '%s'...", versionDir)
	} else if !fs.IsDir() {
		t.Errorf("'%s' must be a folder...", versionDir)
	}

	unpackedFilePath := filepath.Join(versionDir, "bin", "test_file01.txt")
	if fs, err := os.Stat(unpackedFilePath); err != nil {
		t.Errorf("Could not stat unpacked file at '%s'...", unpackedFilePath)
	} else if fs.IsDir() {
		t.Errorf("Go binary '%s' must be a file...", unpackedFilePath)
	}
}