package versions

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eldius/go-version-manager/config"
	"github.com/PuerkitoBio/goquery"
)

/*
ListAvailableVersions lists available Go versions
*/
func ListAvailableVersions() []GoVersion {
	res, err := http.Get("https://golang.org/dl/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return parseDownloadPage(res.Body)

}

func parseDownloadPage(body io.ReadCloser) []GoVersion {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	var versions []GoVersion
	doc.Find("table.codetable").Each(func(i int, t *goquery.Selection) {
		parentAttr, _ := t.Attr("class")
		log.Printf("testing 00: %v (%v/%s)\n", goquery.NodeName(t), t.HasClass("collapsed"), parentAttr)
		t.Parent().Find("h2").Each(func(i int, h *goquery.Selection) {
			version := ParseVersionName(h.Text())
			v := GoVersion{
				Name: version,
			}
			t.Find("tbody>tr").Each(func(i int, r *goquery.Selection) {
				link, _ := r.Find("td.filename>a").Attr("href")
				osName := r.Find("td::nth-child(3)").Text()
				archName := r.Find("td::nth-child(4)").Text()
				log.Printf("os: '%s' / arch: '%s' / link: '%s'", osName, archName, link)
				switch os := fmt.Sprintf("%s-%s", osName, archName); os {
				case "Linux-x86-64":
					v.LinuxAmd64 = parseLink(link)
					log.Println("linux")
				case "-":
					v.Source = parseLink(link)
					log.Println("souce")
				default:

				}
			})
			if v.Source != "" {
				versions = append(versions, v)
			}
		})
	})

	return versions
}

func parseLink(link string) string {
	return fmt.Sprintf("https://golang.org%s", link)
}

/*
ParseVersionName removes the special character
from HTML text
*/
func ParseVersionName(version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(version, " ▹", ""), " ▾", "")
}

/*
Install installs go versions
*/
func Install(version string) {
	v := filterVersion(version, ListAvailableVersions())
	if v == nil {
		fmt.Println("Version not found...")
		return
	}
	f, err := v.Download()
	if err != nil {
		fmt.Printf("Failed to download file from '%s'\n", v.LinuxAmd64)
		log.Panic(err.Error())
	}
	os.Setenv("PATH", fmt.Sprintf("%s:%s", filepath.Join(config.GetWorkspaceDir(), "bin"), os.Getenv("PATH")))
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(f)
}

/*
FilterVersion filter version slice by name
*/
func filterVersion(version string, versions []GoVersion) *GoVersion {
	for _, v := range versions {
		if v.Name == version {
			log.Println(v)
			return &v
		}
	}
	return nil
}
