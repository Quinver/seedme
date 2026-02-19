package search

import (
	"net/http"
	"seedme/internal/model"
	"sort"
	"time"
)

func All(query []string) ([]model.Torrent, error) {
	var results []model.Torrent

	client := &http.Client{
			Timeout: 1 * time.Second,
	}

	nyaa, _ := SearchNyaa(client, query)
	sukebei, _ := SearchSukebei(client, query)
	uIndex, _ := SearchUIndex(client, query)
	pirateBay, _ := SearchPirateBay(client, query)

	results = append(results, nyaa...)
	results = append(results, sukebei...)
	results = append(results, uIndex...)
	results = append(results, pirateBay...)

	sort.Slice(results, func(i, j int) bool {
		return results[i].Seeds > results[j].Seeds
	})

	return results, nil
}
