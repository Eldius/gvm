package versions

import (
	"log"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

const (
	versionDownload = "https://download.com/dl/go1.14.tar.gz"
)

func TestDownload(t *testing.T) {
	defer gock.Off()

	v := GoVersion{
		LinuxAmd64: versionDownload,
	}

	gock.New(versionDownload).
		Get("/").
		Reply(200).
		File("test_data/download.html")

	f, err := v.Download()
	if err != nil {
		t.Error("Failed to download file")
	}

	log.Printf("downloaded file: '%s'", f)
	if f == "" {
		t.Error("File should be not nil")
	}

	i, err := os.Stat(f)
	if err != nil {
		t.Error("Failed validate file state")
	}
	if i.IsDir() {
		t.Error("Returned file path must be a file (but was a directory)")
	}
}
