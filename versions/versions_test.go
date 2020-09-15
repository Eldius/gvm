package versions

import (
	"os"
	"testing"
)

func TestParseDownloadPage(t *testing.T) {
	if body, err := os.Open("test_data/download.html"); err != nil {
		t.Errorf("Failed to read test data")
	} else {
		versions := parseDownloadPage(body)

		if len(versions) == 0 {
			t.Errorf("Versions list is empty")
		}

		t.Logf("---\nlength: %d\n", len(versions))
		for _, v := range versions {
			if v.LinuxAmd64 == "" {
				t.Errorf("Version '%s' must have a LinuxAmd64 link", v.Name)
			}
			if v.Source == "" {
				t.Errorf("Version '%s' must have a Source link", v.Name)
			}
		}
	}
}

func TestFilterVersionValid(t *testing.T) {
	versions := createVersionsSlice()

	v := filterVersion("go1.15.1", versions)

	if v == nil {
		t.Error("v should not be nil")
	} else if v.LinuxAmd64 != "https://link.xpto/1.15.1.tar.gz" {
		t.Errorf("v.LinuxAmd64 should be 'https://link.xpto/1.15.1.tar.gz', but was '%s'", v.LinuxAmd64)
	}
}

func TestFilterVersionInvalid(t *testing.T) {
	versions := createVersionsSlice()

	v := filterVersion("platipus", versions)

	if v != nil {
		t.Error("v should be nil")
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
