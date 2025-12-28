package search

import (
	"fmt"
	"net/http"
	"net/url"
	"seedme/internal/model"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SearchUIndex(query []string) ([]model.Torrent, error) {
	doc, err := fetchUIndex(uIndexURL(query))
	if err != nil {
		panic(err)
	}

	var results []model.Torrent

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(
			s.Find(`a[href^="/details"]`).First().Text(),
			)
		if title == "" {
			return
		}

		magnet, ok := s.Find((`a[href^="magnet:"]`)).Attr("href")
		if !ok {
			return
		}


		seedText := strings.TrimSpace(s.Find("span.g").First().Text())
		seeds, err := strconv.Atoi(seedText)
		if err != nil {
			seeds = 0
		}

		results = append(results, model.Torrent{
			Site:   "uIndex",
			Title:  title,
			Seeds:  seeds,
			Magnet: magnet,
		})
	})

	return results, nil
}

func uIndexURL(query []string) string {
	q := url.QueryEscape(strings.Join(query, " "))
	return fmt.Sprintf(
		"https://uindex.org/search.php?search=%s&c=0&sort=seeders&order=DESC",
		q,
	)
}

func fetchUIndex(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Status code not OK")
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
