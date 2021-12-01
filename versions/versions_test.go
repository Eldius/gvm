package versions

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
	"gopkg.in/h2non/gock.v1"
)

const (
	downloadListPage = "https://link.xpto/dl"
)

func init() {
	viper.SetDefault("gvm.versions.page.url", downloadListPage)
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("falha carregando arquivo de configuracoes\n", err.Error())
	}

}

func TestParseDownloadPage(t *testing.T) {
	if body, err := os.Open("test_data/download.html"); err != nil {
		t.Errorf("Failed to read test data")
	} else {
		versions := parseDownloadPage(body)

		if len(versions) == 0 {
			t.Errorf("Versions list is empty")
		}

		v1_16_5 := false
		t.Logf("---\nlength: %d\n", len(versions))
		for _, v := range versions {
			if v.Name == "go1.16.5" {
				v1_16_5 = true
			}
			if v.LinuxAmd64 == "" {
				t.Errorf("Version '%s' must have a LinuxAmd64 link", v.Name)
			}
			if v.Source == "" {
				t.Errorf("Version '%s' must have a Source link", v.Name)
			}
		}

		if !v1_16_5 {
			t.Error("Must find version 'go1.16.5'")
		}
	}
}

func TestFilterVersionValid(t *testing.T) {
	versions := createVersionsSlice()

	v := FilterVersion("go1.15.1", versions)

	if v == nil {
		t.Error("v should not be nil")
	} else if v.LinuxAmd64 != "https://link.xpto/1.15.1.tar.gz" {
		t.Errorf("v.LinuxAmd64 should be 'https://link.xpto/1.15.1.tar.gz', but was '%s'", v.LinuxAmd64)
	}
}

func TestFilterVersionValidWithoutGo(t *testing.T) {
	versions := createVersionsSlice()

	v := FilterVersion("1.15.1", versions)

	if v == nil {
		t.Error("v should not be nil")
	} else if v.LinuxAmd64 != "https://link.xpto/1.15.1.tar.gz" {
		t.Errorf("v.LinuxAmd64 should be 'https://link.xpto/1.15.1.tar.gz', but was '%s'", v.LinuxAmd64)
	}
}

func TestFilterVersionInvalid(t *testing.T) {
	versions := createVersionsSlice()

	v := FilterVersion("platipus", versions)

	if v != nil {
		t.Error("v should be nil")
	}
}

func TestListAvailableVersions(t *testing.T) {
	qtdVersions := 201
	defer gock.Off()

	gock.New(downloadListPage).
		Get("/").
		Reply(200).
		File("test_data/download.html")

	versions := ListAvailableVersions()

	if len(versions) != qtdVersions {
		t.Errorf("Should have %d versions but has '%d'", qtdVersions, len(versions))
	}
}

func createVersionsSlice() []GoVersion {
	return []GoVersion{
		{
			Name:       "go1.15.1",
			LinuxAmd64: "https://link.xpto/1.15.1.tar.gz",
		}, {
			Name:       "go1.14.8",
			LinuxAmd64: "https://link.xpto/1.14.8.tar.gz",
		},
	}
}
