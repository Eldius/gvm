package updates

import (
	"os"
	"testing"
)

func TestParseReleasesResponse(t *testing.T) {

	f, err := os.Open("test_data/sample.json")
	if err != nil {
		t.Errorf("Failed to open sample file: %s", err.Error())
	}

	r := parseReleasesResponse(f)

	if len(r) != 3 {
		t.Errorf("Releases list lenth must have 3 elements, but has %d", len(r))
	}
}
