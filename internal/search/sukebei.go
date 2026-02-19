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

func SearchSukebei(client *http.Client, query []string) ([]model.Torrent, error) {
	doc, err := fetchSukebei(client, sukebeiURL(query))
	if err != nil {
		return nil, fmt.Errorf("Fetch failed: %w", err)
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
			Site:   "Sukebei",
			Title:  title,
			Seeds:  seeds,
			Magnet: magnet,
		})
	})

	return results, nil
}

func sukebeiURL(query []string) string {
	q := url.QueryEscape(strings.Join(query, " "))
	return fmt.Sprintf(
		"https://sukebei.nyaa.si/?f=0&c=0_0&q=%s&s=seeders&o=desc",
		q,
	)
}

func fetchSukebei(client *http.Client, url string) (*goquery.Document, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Can't get response: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
