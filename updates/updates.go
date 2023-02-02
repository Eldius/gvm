/*
Package updates has the gvm version management code
(`experimental`)
*/
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
- {{ .GetName }}:
    published: {{ .GetPublishedAt }}
    assets:{{ range .GetArtifacts }}
      - {{ .GetName }}
        download: {{ .GetArtifactURL }}{{ end }}`
)

/*
CheckForUpdates check for gvm updates at GitHub packages
*/
func CheckForUpdates() {
    suffixPattern := fmt.Sprintf("%s.%s", runtime.GOOS, runtime.GOARCH)
    fmt.Printf("Architecture: %s", suffixPattern)
    r, err := updater.ListReleases("eldius", "gvm", "GITHUB")
    if err != nil {
        fmt.Printf("Failed to get releases: %s", err.Error())
    }

    for _, v := range r {
        t := template.Must(template.New("main").Parse(versiomTemplate))
        _ = t.Execute(os.Stdout, v)
    }
    fmt.Print("\n\n")
}
