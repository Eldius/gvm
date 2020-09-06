package versions

import (
	"fmt"
	"log"
	"net/http"
	"strings"

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

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var versions []GoVersion
	// Find the review items
	doc.Find("h2.toggleButton[title=\"Click to show downloads for this version\"]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		version := ParseVersionName(s.Text())
		fmt.Printf("version %d: '%s'\n", i, version)
		s.Parent().Find("table.codetable>tbody>a.download").Each(func(i1 int, l *goquery.Selection) {
			log.Println("testing...")
			log.Printf("OS: %s\n", l.Find("..>..>td(2)").Text())
			log.Printf("Arch: %s\n", l.Find("..>..>td(3)").Text())
		})
		versions = append(versions, GoVersion{
			Name: version,
		})
	})

	return versions

}

/*
ParseVersionName removes the special character
from HTML text
*/
func ParseVersionName(version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(version, " ▹", ""), " ▾", "")
}
