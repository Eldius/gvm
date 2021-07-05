package updates

import (
	"fmt"
	"os"
	"runtime"
	"text/template"

	"github.com/Eldius/app-releases-go/updater"
)

const (
	versiomTemplate = `
---
- {{ .Name }}:
  published: {{ .PublishedAt }}
  created:   {{ .CreatedAt }}
  assets:{{ range .Assets }}
    - {{ .Name }}
      download: {{ .BrowserDownloadURL }}{{ end }}`
)

func CheckForUpdates() {
	fmt.Printf("Architecture: %s.%s", runtime.GOOS, runtime.GOARCH)
	r, err := updater.ListReleases("eldius", "gvm", "GITHUB")
	if err != nil {
		fmt.Printf("Failed to get releases: %s", err.Error())
	}

	for _, v := range r {
		t := template.Must(template.New("main").Parse(versiomTemplate))
		t.Execute(os.Stdout, v)
	}
	fmt.Print("\n\n")
}
