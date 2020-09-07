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
