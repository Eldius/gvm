package versions

import (
	"os"
	"testing"
)

func TestParseDownloadPage(t *testing.T) {
	if body, err := os.Open("test_data/download.html"); err != nil {
		t.Errorf("Failed to read test data")

		versions := parseDownloadPage(body)

		if len(versions) == 0 {
			t.Errorf("Versions list is empty")
		}

		t.Logf("---\nlength: %d\n", len(versions))
		for _, v := range versions {
			t.Logf("---\nversion: %s\n", v.Name)
		}
	}
}