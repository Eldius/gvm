//go:build linux && arm64

package versions

func TestGoVersion_GetURL(t *testing.T) {
	if body, err := os.Open("test_data/download.html"); err != nil {
		t.Errorf("Failed to read test data")
	} else {
		versions := parseDownloadPage(body)

		if len(versions) == 0 {
			t.Errorf("Versions list is empty")
		}
		v1_20_4 := false
		t.Logf("---\nlength: %d\n", len(versions))
		for _, v := range versions {
			if v.Name == "go1.20.4" {
				v1_20_4 = true
				assert.True(t, strings.HasSuffix(v.GetURL(), "go1.20.4.linux-arm64.tar.gz"))
			}
		}

		v1_16_5 := false
		for _, v := range versions {
			if v.Name == "go1.16.5" {
				v1_16_5 = true
				assert.True(t, strings.HasSuffix(v.GetURL(), "go1.16.5.linux-arm64.tar.gz"))
			}
		}

		if !v1_16_5 {
			t.Error("Must find version 'go1.16.5'")
		}

		if !v1_20_4 {
			t.Error("Must find version 'go1.20.4'")
		}
	}
}
