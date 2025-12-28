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

func SearchEZTV(query []string) ([]model.Torrent, error) {
	doc, err := fetchEZTV(eztvURL(query))
	if err != nil {
		panic(err)
	}

	var results []model.Torrent

	doc.Find("table tbody tr").Each(func(i int, s *goquery.Selection) {

		a := s.Find("a.epinfo").First()
		if a.Length() == 0 {
			return
		}

		href, ok := a.Attr("href")
		if !ok {
			return
		}

		title, _ := a.Attr("href")
		epURL := "https://eztvx.to" + href
		
		seedText := strings.TrimSpace(
			s.Find("td.forum_thread_post_end").First().Text(),
		)

		seeds, _ := strconv.Atoi(seedText)

		magnet, err := fetchEZTVMagnet(epURL)
		if err != nil {
			return
		}

		results = append(results, model.Torrent{
			Site:   "eztv",
			Title:  title,
			Seeds:  seeds,
			Magnet: magnet,
		})
	})

	return results, nil
}

func eztvURL(query []string) string {
	q := url.QueryEscape(strings.Join(query, " "))
	return fmt.Sprintf(
		"https://eztvx.to/search/%s",
		q,
	)
}

func fetchEZTV(url string) (*goquery.Document, error) {
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

func fetchEZTVMagnet(epURL string) (string, error) {
	resp, err := http.Get(epURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	magnet, ok := doc.
		Find(`a[title="Magnet Link"]`).
		Attr("href")

	if !ok {
		return "", fmt.Errorf("no magnet link found")
	}

	return magnet, nil
}
