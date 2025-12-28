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

func SearchNyaa(query []string) ([]model.Torrent, error) {
	doc, err := fetchNyaa(nyaaURL(query))
	if err != nil {
		panic(err)
	}

	var results []model.Torrent

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {
		a := s.Find(`a[href^="/view/"]:not(.comments)`).First()
		title, ok := a.Attr("title")
		if !ok {
			return
		}

		magnet, ok := s.Find((`a[href^="magnet:"]`)).Attr("href")
		if !ok {
			return
		}

		tds := s.Find("td.text-center")
		if tds.Length() < 4 {
			return
		}

		seedText := strings.TrimSpace(tds.Eq(3).Text())
		seeds, err := strconv.Atoi(seedText)
		if err != nil {
			seeds = 0
		}

		results = append(results, model.Torrent{
			Site:   "nyaa",
			Title:  title,
			Seeds:  seeds,
			Magnet: magnet,
		})
	})

	return results, nil
}

func nyaaURL(query []string) string {
	q := url.QueryEscape(strings.Join(query, " "))
	return fmt.Sprintf(
		"https://nyaa.si/?f=0&c=0_0&q=%s&s=seeders&o=desc",
		q,
	)
}

func fetchNyaa(url string) (*goquery.Document, error) {
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
