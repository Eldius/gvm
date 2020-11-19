package versions

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Eldius/gvm/config"
	"github.com/PuerkitoBio/goquery"
)

/*
ListAvailableVersions lists available Go versions
*/
func ListAvailableVersions() []GoVersion {
	res, err := http.Get(config.GetVersionsPage())
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
FilterVersion filter version slice by name
*/
func FilterVersion(version string, versions []GoVersion) *GoVersion {
	for _, v := range versions {
		if CompareVersions(version, v.Name) {
			log.Println(v)
			return &v
		}
	}
	return nil
}

/*
CompareVersions compares a version with the required version
*/
func CompareVersions(required string, version string) bool {
	return version == required || strings.Replace(version, "go", "", 1) == required
}

func ListLocalVersions() (result []string) {
	files, err := ioutil.ReadDir(config.GetVersionsDir())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		result = append(result, f.Name())
	}

	return result
}
