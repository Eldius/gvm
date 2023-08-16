package versions

import (
	"log"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

const (
	versionDownloadLink     = "https://download.com/dl/go1.14.tar.gz"
	versionDownloadChecksum = "238c0d2cf99b0ee780f756d68f86e4c711f757af6109db767f82e2edc93dfd00"
)

func TestDownload(t *testing.T) {
	defer gock.Off()

	v := GoVersion{
		Name:               "go1.14",
		LinuxAmd64:         versionDownloadLink,
		LinuxAmd64Checksum: versionDownloadChecksum,
	}

	gock.New(versionDownloadLink).
		Get("/").
		Reply(200).
		File("test_data/download.html")

	f, err := v.Download()
	if err != nil {
		t.Error("Failed to download file:", err)
		t.FailNow()
	}

	log.Printf("downloaded file: '%s'", f)
	if f == "" {
		t.Error("File should be not nil:", err)
		t.FailNow()
	}

	i, err := os.Stat(f)
	if err != nil {
		t.Error("Failed validate file state:", err)
		t.FailNow()
	}
	if i.IsDir() {
		t.Error("Returned file path must be a file (but was a directory):", err)
		t.FailNow()
	}
}
