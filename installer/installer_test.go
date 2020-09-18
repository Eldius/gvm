package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	log.Println("temp dir:", tmpDir)
	os.Setenv("USER", tmpDir)
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

}