package updates

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
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
	res, err := http.Get("https://api.github.com/repos/eldius/gvm/releases")
	if err != nil {
		log.Fatalf("Failed to fetch app releases: %s\n", err.Error())
	}
	defer res.Body.Close()

	r := parseReleasesResponse(res.Body)

	for _, v := range r {
		t := template.Must(template.New("main").Parse(versiomTemplate))
		t.Execute(os.Stdout, v)
	}
	fmt.Print("\n\n")

}

func parseReleasesResponse(body io.ReadCloser) []*GithubReleasesResponse {
	var response []*GithubReleasesResponse
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		log.Fatalf("Failed to parse releases response: %s\n", err.Error())
	}

	return response
}
